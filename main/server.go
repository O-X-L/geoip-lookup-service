package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/superstes/geoip-lookup-service/cnf"
	"github.com/superstes/geoip-lookup-service/lookup"
	"github.com/superstes/geoip-lookup-service/u"
)

func errorResponse(w http.ResponseWriter, m string) {
	w.WriteHeader(http.StatusBadRequest)
	_, err := io.WriteString(w, fmt.Sprintf("%v\n", m))
	if err != nil {
		log.Fatal(err)
	}
}

func returnResult(w http.ResponseWriter, data interface{}) {
	if cnf.RETURN_PLAIN {
		w.Header().Set("Content-Type", "text/plain")
		_, err := io.WriteString(w, fmt.Sprintf("%+v\n", data))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Fatal(err)
			errorResponse(w, "Failed to JSON-encode data!")
		}
	}
}

func geoIpLookup(w http.ResponseWriter, r *http.Request) {
	ipStr := r.URL.Query().Get("ip")
	lookupStr := r.URL.Query().Get("lookup")
	filterStr := r.URL.Query().Get("filter")
	if lookupStr == "" || ipStr == "" {
		errorResponse(w, "Either 'lookup' or 'ip' were not provided!")
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
		log.Fatal(err)
		errorResponse(w, "Failed to lookup data")
		return
	}

	if filterStr != "" {
		filteredData := data
		for _, subFilterStr := range strings.Split(filterStr, ".") {
			defer func() {
				if err := recover(); err != nil {
					log.Fatal(err)
					errorResponse(w, "Invalid FILTER provided")
				}
			}()
			filteredData = u.GetMapValue(filteredData, subFilterStr)
			if filteredData == nil {
				errorResponse(w, "Invalid FILTER provided")
				return
			}
		}
		returnResult(w, filteredData)
		return
	}

	returnResult(w, data)
	return
}

func httpServer(listenAddr string, listenPort uint) {
	http.HandleFunc("/", geoIpLookup)
	var listenStr = fmt.Sprintf("%v:%v", listenAddr, listenPort)
	fmt.Println("Listening on http://" + listenStr)
	log.Fatal(http.ListenAndServe(listenStr, nil))
}
