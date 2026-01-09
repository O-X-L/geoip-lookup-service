package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"

	"git.oxl.at/geoip-lookup-service/internal/cnf"
	"git.oxl.at/geoip-lookup-service/internal/lookup"
	u "git.oxl.at/geoip-lookup-service/internal/util"
)

func errorResponse(w http.ResponseWriter, m string) {
	w.WriteHeader(http.StatusBadRequest)
	_, err := io.WriteString(w, fmt.Sprintf("%v\n", m))
	if err != nil {
		log.Fatal(err)
	}
}

func returnResult(w http.ResponseWriter, data interface{}, logPrefix string) {
	result := fmt.Sprintf("%+v\n", data)
	if cnf.RETURN_PLAIN || (!strings.Contains(result, "[") && !strings.Contains(result, "{")) {
		w.Header().Set("Content-Type", "text/plain")
		_, err := io.WriteString(w, result)
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

func handleGeoIPLookup(w http.ResponseWriter, r *http.Request) {
	ipStr := r.URL.Query().Get("ip")
	lookupStr := r.URL.Query().Get("lookup")
	filterStr := r.URL.Query().Get("filter")
	logPrefix := fmt.Sprintf("lookup=%v, filter=%v, ip=%v", lookupStr, filterStr, ipStr)

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

	// ease-of-use: allow users to only supply commonly used filters - we pick the correct DB-type for them
	switch filterStr {
	case "":
		if value, exists := cnf.LOOKUP_FILTER_SHORTCUTS[cnf.DB_TYPE][lookupStr]; exists {
			filterStr = lookupStr
			lookupStr = value
		}
	case "*":
		filterStr = ""
	}
	logPrefix = fmt.Sprintf("lookup=%v, filter=%v, ip=%v", lookupStr, filterStr, ipStr)

	lookupFunc, validLookup := lookup.FUNC_MAPPING[cnf.DB_TYPE][lookupStr]
	if lookupFunc == nil || !validLookup {
		u.LogWarn(logPrefix, "Invalid Lookup")
		errorResponse(w, "Invalid LOOKUP provided")
		return
	}

	ip := net.ParseIP(ipStr)
	if ip == nil {
		u.LogWarn(logPrefix, "Invalid IP")
		errorResponse(w, "Invalid IP provided")
		return
	}
	logPrefix = fmt.Sprintf("lookup=%v, filter=%v, ip=%v", lookupStr, filterStr, ipStr)

	data, err := lookupFunc(ip)
	if data == nil {
		u.LogWarn(logPrefix, "Invalid Lookup")
		errorResponse(w, "Invalid LOOKUP provided")
		return
	}
	if err != nil {
		u.LogError(logPrefix, fmt.Sprintf("Lookup failed: %v", err))
		errorResponse(w, "Failed to lookup data")
		return
	}

	if filterStr == "" {
		returnResult(w, data, logPrefix)
		return
	}

	defer func() {
		if err := recover(); err != nil {
			// private ips or non-existant attributes
			u.LogWarn(logPrefix, "IP not in MMDB or filtering on non-existant attribute")
			returnResult(w, "", logPrefix)
		}
	}()
	filteredData := data
	for _, subFilterStr := range strings.Split(filterStr, ".") {
		filteredData = u.GetMapValue(filteredData, subFilterStr)
		if filteredData == nil {
			u.LogWarn(logPrefix, "Filtering on non-existant attribute")
			returnResult(w, "", logPrefix)
			return
		}
	}
	returnResult(w, filteredData, logPrefix)
}

func HttpServer(listenAddr string, listenPort uint) {
	http.HandleFunc("/", handleGeoIPLookup)
	var listenStr = fmt.Sprintf("%v:%v", listenAddr, listenPort)
	fmt.Println("Listening on http://" + listenStr)
	log.Fatal(http.ListenAndServe(listenStr, nil))
}
