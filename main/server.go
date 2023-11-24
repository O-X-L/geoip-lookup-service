package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"

	"github.com/superstes/geoip-lookup-service/cnf"
	"github.com/superstes/geoip-lookup-service/lookup"
	"github.com/superstes/geoip-lookup-service/u"
)

func errorResponse(w http.ResponseWriter, m string) {
	w.WriteHeader(http.StatusBadRequest)
	_, err := io.WriteString(w, m)
	if err != nil {
		log.Fatal(err)
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
		errorResponse(w, "Provided IP is not valid!")
		return
	}

	dataStructure, lookupExists := cnf.LOOKUP[lookupStr]
	if !lookupExists || dataStructure == nil {
		errorResponse(w, "Provided LOOKUP is not valid!")
		return
	}

	data, err := lookup.FUNC[lookupStr].(func(net.IP, interface{}) (interface{}, error))(
		ip, dataStructure,
	)
	if err != nil {
		log.Fatal(err)
		errorResponse(w, "Failed to lookup data!")
		return
	}

	if filterStr != "" {
		// todo: allow deeper filtering
		defer func() {
			if err := recover(); err != nil {
				log.Fatal(err)
				errorResponse(w, "Provided FILTER is not valid!")
			}
		}()
		filteredData := u.GetAttribute(data, filterStr)
		if !filteredData.IsValid() {
			errorResponse(w, "Provided FILTER is not valid!")
		}
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(filteredData.String())
		if err != nil {
			log.Fatal(err)
			errorResponse(w, "Failed to JSON-encode data!")
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatal(err)
		errorResponse(w, "Failed to JSON-encode data!")
	}
	return
}

func httpServer(listenAddr string, listenPort uint) {
	http.HandleFunc("/", geoIpLookup)
	var listenStr = fmt.Sprintf("%v:%v", listenAddr, listenPort)
	fmt.Println("Listening on http://" + listenStr)
	log.Fatal(http.ListenAndServe(listenStr, nil))
}
