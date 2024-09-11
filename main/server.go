package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/O-X-L/geoip-lookup-service/cnf"
	"github.com/O-X-L/geoip-lookup-service/lookup"
	"github.com/O-X-L/geoip-lookup-service/u"
)

func errorResponse(w http.ResponseWriter, m string) {
	w.WriteHeader(http.StatusBadRequest)
	_, err := io.WriteString(w, fmt.Sprintf("%v\n", m))
	if err != nil {
		log.Fatal(err)
	}
}

func returnResult(w http.ResponseWriter, data interface{}, logPrefix string) {
	if cnf.RETURN_PLAIN {
		w.Header().Set("Content-Type", "text/plain")
		_, err := io.WriteString(w, fmt.Sprintf("%+v\n", data))
		if err != nil {
			u.LogError(logPrefix, err)
		}
	} else {
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			u.LogError(logPrefix, err)
			errorResponse(w, "Failed to JSON-encode data")
		}
	}
}

func getClientIP(r *http.Request) (string, error) {
	fwdIPs := strings.Split(r.Header.Get("X-Forwarded-For"), ",")
	if len(fwdIPs) > 0 {
		netIP := net.ParseIP(fwdIPs[len(fwdIPs)-1])
		if netIP != nil {
			return netIP.String(), nil
		}
	}

	realIP := r.Header.Get("X-Real-IP")
	if realIP != "" {
		netIP := net.ParseIP(realIP)
		if netIP != nil {
			return netIP.String(), nil
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	netIP := net.ParseIP(ip)
	if netIP != nil {
		ip := netIP.String()
		if ip == "::1" {
			return "127.0.0.1", nil
		}
		return ip, nil
	}

	return "", errors.New("IP not found")
}

func geoIpLookup(w http.ResponseWriter, r *http.Request) {
	ipStr := r.URL.Query().Get("ip")
	lookupStr := r.URL.Query().Get("lookup")
	filterStr := r.URL.Query().Get("filter")
	logPrefix := fmt.Sprintf("IP: '%v', Lookup: '%v', Filter: '%v'", ipStr, lookupStr, filterStr)

	if ipStr == "" {
		clientIpStr, err := getClientIP(r)
		if err == nil {
			ipStr = clientIpStr
		}
	}

	if lookupStr == "" || ipStr == "" {
		errorResponse(w, "Either 'lookup' or 'ip' were not provided")
		return
	}

	ip := net.ParseIP(ipStr)
	if ip == nil {
		errorResponse(w, "Invalid IP provided")
		return
	}

	data, err := lookup.FUNC[lookupStr].(func(net.IP) (interface{}, error))(ip)
	if data == nil {
		errorResponse(w, "Invalid LOOKUP provided")
		return
	}
	if err != nil {
		u.LogError(logPrefix, err)
		errorResponse(w, "Failed to lookup data")
		return
	}

	if filterStr != "" {
		filteredData := data
		for _, subFilterStr := range strings.Split(filterStr, ".") {
			defer func() {
				if err := recover(); err != nil {
					u.LogError(logPrefix, err)
					errorResponse(w, "Invalid FILTER provided")
				}
			}()
			filteredData = u.GetMapValue(filteredData, subFilterStr)
			if filteredData == nil {
				errorResponse(w, "Invalid FILTER provided")
				return
			}
		}
		returnResult(w, filteredData, logPrefix)
		return
	}

	returnResult(w, data, logPrefix)
}

func httpServer(listenAddr string, listenPort uint) {
	http.HandleFunc("/", geoIpLookup)
	var listenStr = fmt.Sprintf("%v:%v", listenAddr, listenPort)
	fmt.Println("Listening on http://" + listenStr)
	log.Fatal(http.ListenAndServe(listenStr, nil))
}
