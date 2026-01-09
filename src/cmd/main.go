package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"git.oxl.at/geoip-lookup-service/internal"
	"git.oxl.at/geoip-lookup-service/internal/cnf"
	u "git.oxl.at/geoip-lookup-service/internal/util"
)

func welcome() {
	fmt.Printf("\n   ______           ________     __                __             \n")
	fmt.Println("  / ____/__  ____  /  _/ __ \\   / /   ____  ____  / /____  ______ ")
	fmt.Println(" / / __/ _ \\/ __ \\ / // /_/ /  / /   / __ \\/ __ \\/ //_/ / / / __ \\")
	fmt.Println("/ /_/ /  __/ /_/ // // ____/  / /___/ /_/ / /_/ / ,< / /_/ / /_/ /")
	fmt.Println("\\____/\\___/\\____/___/_/      /_____/\\____/\\____/_/|_|\\__,_/ .___/ ")
	fmt.Println("                                                         /_/      ")
	fmt.Printf("Version: %s\n", cnf.VERSION)
	fmt.Printf("by OXL IT Services (License: MIT)\n\n")
}

func main() {
	var listenAddr string
	var listenPort uint
	var dbType string

	flag.StringVar(&listenAddr, "l", "127.0.0.1", "Address to listen on")
	flag.UintVar(&listenPort, "p", 10000, "Port to listen on")
	flag.StringVar(&dbType, "t", "ipinfo", "Database type to use (ipinfo or maxmind)")
	flag.StringVar(&cnf.DB_LITE, "lite", cnf.DB_LITE, "Path to the Lite-database (only for IPInfo; optional)")
	flag.StringVar(&cnf.DB_COUNTRY, "country", cnf.DB_COUNTRY, "Path to the country-database (optional)")
	flag.StringVar(&cnf.DB_CITY, "city", cnf.DB_CITY, "Path to the city-database (optional)")
	flag.StringVar(&cnf.DB_ASN, "asn", cnf.DB_ASN, "Path to the asn-database (optional)")
	flag.StringVar(&cnf.DB_PRIVACY, "privacy", cnf.DB_PRIVACY, "Path to the privacy-database (optional)")
	flag.BoolVar(&cnf.RETURN_PLAIN, "plain", cnf.RETURN_PLAIN, "If the result should be returned in plain text format")
	flag.BoolVar(&cnf.CLIENT_IP_FWD_HDR, "ip-fwd-hdr", cnf.CLIENT_IP_FWD_HDR, "If no IP was provided - try to pull the client-IP from the Forwarded-For header")
	flag.Parse()

	dbType = strings.ToLower(dbType)

	switch dbType {
	case "maxmind":
		cnf.DB_TYPE = cnf.DB_TYPE_MAXMIND
	case "oxl":
		cnf.DB_TYPE = cnf.DB_TYPE_OXL
	default:
		cnf.DB_TYPE = cnf.DB_TYPE_IPINFO
	}

	welcome()

	u.PrependDBDirPaths()
	if err := u.CheckGeoIPDBs(); err != nil {
		u.LogError("", err)
		os.Exit(1)
	}

	internal.HttpServer(listenAddr, listenPort)
}
