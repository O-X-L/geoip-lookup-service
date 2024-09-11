package main

import (
	"flag"
	"fmt"

	"github.com/O-X-L/geoip-lookup-service/cnf"
)

func welcome() {
	fmt.Printf("\n   ______           ________     __                __             \n")
	fmt.Println("  / ____/__  ____  /  _/ __ \\   / /   ____  ____  / /____  ______ ")
	fmt.Println(" / / __/ _ \\/ __ \\ / // /_/ /  / /   / __ \\/ __ \\/ //_/ / / / __ \\")
	fmt.Println("/ /_/ /  __/ /_/ // // ____/  / /___/ /_/ / /_/ / ,< / /_/ / /_/ /")
	fmt.Println("\\____/\\___/\\____/___/_/      /_____/\\____/\\____/_/|_|\\__,_/ .___/ ")
	fmt.Println("                                                         /_/      ")
	fmt.Printf("Version: %v\n", cnf.VERSION)
	fmt.Printf("by OXL IT Services (License: MIT)\n\n")
}

func main() {
	var listenAddr string
	var listenPort uint
	var dbType string

	flag.StringVar(&listenAddr, "l", "127.0.0.1", "Address to listen on")
	flag.UintVar(&listenPort, "p", 10000, "Port to listen on")
	flag.StringVar(&dbType, "t", "ipinfo", "Database type to use (ipinfo or maxmind)")
	flag.StringVar(&cnf.DB_COUNTRY, "country", cnf.DB_COUNTRY, "Path to the country-database (optional)")
	flag.StringVar(&cnf.DB_CITY, "city", cnf.DB_CITY, "Path to the city-database (optional)")
	flag.StringVar(&cnf.DB_ASN, "asn", cnf.DB_ASN, "Path to the asn-database (optional)")
	flag.StringVar(&cnf.DB_PRIVACY, "privacy", cnf.DB_PRIVACY, "Path to the privacy-database (optional)")
	flag.BoolVar(&cnf.RETURN_PLAIN, "plain", cnf.RETURN_PLAIN, "If the result should be returned in plain text format")
	flag.Parse()

	if dbType == "maxmind" {
		cnf.DB_TYPE = cnf.DB_TYPE_MAXMIND
	} else {
		cnf.DB_TYPE = cnf.DB_TYPE_IPINFO
	}

	welcome()
	httpServer(listenAddr, listenPort)
}
