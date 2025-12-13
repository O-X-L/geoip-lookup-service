package lookup

import (
	"net"

	"git.oxl.at/geoip-lookup-service/internal/cnf"
)

func IpInfoCountry(ip net.IP) (interface{}, error) {
	if cnf.DB_LITE != "" {
		return lookupBase(ip, cnf.IPINFO_LITE, cnf.DB_LITE)
	}
	return lookupBase(ip, cnf.IPINFO_COUNTRY, cnf.DB_COUNTRY)
}

func IpInfoAsn(ip net.IP) (interface{}, error) {
	if cnf.DB_LITE != "" {
		return lookupBase(ip, cnf.IPINFO_LITE, cnf.DB_LITE)
	}
	return lookupBase(ip, cnf.IPINFO_ASN, cnf.DB_ASN)
}

func IpInfoCity(ip net.IP) (interface{}, error) {
	return lookupBase(ip, cnf.IPINFO_CITY, cnf.DB_CITY)
}

func IpInfoCountryAsn(ip net.IP) (interface{}, error) {
	return lookupBase(ip, cnf.IPINFO_COUNTRY_ASN, cnf.DB_COUNTRY)
}

func IpInfoPrivacy(ip net.IP) (interface{}, error) {
	return lookupBase(ip, cnf.IPINFO_PRIVACY, cnf.DB_PRIVACY)
}
