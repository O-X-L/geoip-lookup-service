package cnf

import "net"

const DB_TYPE_IPINFO uint = 1

// IPInfo schema: https://github.com/ipinfo/sample-database/
type IPINFO_LITE struct {
	Network       net.IP `maxminddb:"network"`
	Country       string `maxminddb:"country_code"`
	CountryName   string `maxminddb:"country"`
	Continent     string `maxminddb:"continent_code"`
	ContinentName string `maxminddb:"continent"`
	ASN           string `maxminddb:"asn"`
	Name          string `maxminddb:"as_name"`
	Domain        string `maxminddb:"as_domain"`
}

type IPINFO_COUNTRY struct {
	StartIp       net.IP `maxminddb:"start_ip"`
	EndIp         net.IP `maxminddb:"end_ip"`
	Country       string `maxminddb:"country"`
	CountryName   string `maxminddb:"country_name"`
	Continent     string `maxminddb:"continent"`
	ContinentName string `maxminddb:"continent_name"`
}

type IPINFO_ASN struct {
	StartIp net.IP `maxminddb:"start_ip"`
	EndIp   net.IP `maxminddb:"end_ip"`
	ASN     string `maxminddb:"asn"`
	Name    string `maxminddb:"name"`
	Domain  string `maxminddb:"domain"`
}

type IPINFO_ASN_EXT struct {
	StartIp net.IP `maxminddb:"start_ip"`
	EndIp   net.IP `maxminddb:"end_ip"`
	JoinKey net.IP `maxminddb:"join_key"`
	ASN     string `maxminddb:"asn"`
	Name    string `maxminddb:"name"`
	Domain  string `maxminddb:"domain"`
	Type    string `maxminddb:"type"`
	Country string `maxminddb:"country"`
}

type IPINFO_COUNTRY_ASN struct {
	StartIp       net.IP `maxminddb:"start_ip"`
	EndIp         net.IP `maxminddb:"end_ip"`
	Country       string `maxminddb:"country"`
	CountryName   string `maxminddb:"country_name"`
	Continent     string `maxminddb:"continent"`
	ContinentName string `maxminddb:"continent_name"`
	ASN           string `maxminddb:"asn"`
	ASName        string `maxminddb:"as_name"`
	ASDomain      string `maxminddb:"as_domain"`
}

type IPINFO_PRIVACY struct {
	StartIp net.IP `maxminddb:"start_ip"`
	EndIp   net.IP `maxminddb:"end_ip"`
	JoinKey net.IP `maxminddb:"join_key"`
	Hosting bool   `maxminddb:"hosting"`
	Proxy   bool   `maxminddb:"proxy"`
	Tor     bool   `maxminddb:"tor"`
	Vpn     bool   `maxminddb:"vpn"`
	Relay   bool   `maxminddb:"relay"`
	Service string `maxminddb:"service"`
}

type IPINFO_CITY struct {
	StartIp    net.IP  `maxminddb:"start_ip"`
	EndIp      net.IP  `maxminddb:"end_ip"`
	JoinKey    net.IP  `maxminddb:"join_key"`
	City       string  `maxminddb:"city"`
	Region     string  `maxminddb:"region"`
	Country    string  `maxminddb:"country"`
	Latitude   float32 `maxminddb:"latitude"`
	Longitude  float32 `maxminddb:"longitude"`
	PostalCode string  `maxminddb:"postal_code"`
	Timezone   string  `maxminddb:"timezone"`
}

// todo: https://github.com/ipinfo/sample-database/tree/main/IP%20to%20Company
// todo: https://github.com/ipinfo/sample-database/tree/main/IP%20to%20Mobile%20Carrier
// todo: https://github.com/ipinfo/sample-database/tree/main/IP%20Geolocation%20Extended
// todo: https://github.com/ipinfo/sample-database/tree/main/Privacy%20Detection%20Extended
// todo: https://github.com/ipinfo/sample-database/tree/main/Abuse%20Contact
// todo: https://github.com/ipinfo/sample-database/tree/main/Hosted%20Domains
