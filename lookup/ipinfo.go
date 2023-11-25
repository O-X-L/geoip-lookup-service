package lookup

import (
	"net"

	"github.com/superstes/geoip-lookup-service/cnf"
)

func IpInfoCountry(ip net.IP) (interface{}, error) {
	return lookupBase(ip, cnf.IPINFO_COUNTRY, cnf.DB_COUNTRY)
}

func IpInfoCity(ip net.IP) (interface{}, error) {
	return lookupBase(ip, cnf.IPINFO_CITY, cnf.DB_CITY)
}

func IpInfoAsn(ip net.IP) (interface{}, error) {
	return lookupBase(ip, cnf.IPINFO_ASN, cnf.DB_ASN)
}

func IpInfoCountryAsn(ip net.IP) (interface{}, error) {
	return lookupBase(ip, cnf.IPINFO_COUNTRY_ASN, cnf.DB_COUNTRY)
}

func IpInfoPrivacy(ip net.IP) (interface{}, error) {
	return lookupBase(ip, cnf.IPINFO_PRIVACY, cnf.DB_PRIVACY)
}
