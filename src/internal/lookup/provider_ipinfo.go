package lookup

import (
	"net"

	"git.oxl.at/geoip-lookup-service/internal/cnf"
)

func IpInfoCountry(ip net.IP) (interface{}, error) {
	if cnf.DB_LITE != "" {
		data := cnf.IPINFO_LITE{}
		return lookupBase(ip, data, cnf.DB_LITE)
	}
	data := cnf.IPINFO_COUNTRY{}
	return lookupBase(ip, data, cnf.DB_COUNTRY)
}

func IpInfoAsn(ip net.IP) (interface{}, error) {
	if cnf.DB_LITE != "" {
		data := cnf.IPINFO_LITE{}
		return lookupBase(ip, data, cnf.DB_LITE)
	}
	data := cnf.IPINFO_ASN{}
	return lookupBase(ip, data, cnf.DB_ASN)
}

func IpInfoCity(ip net.IP) (interface{}, error) {
	data := cnf.IPINFO_CITY{}
	return lookupBase(ip, data, cnf.DB_CITY)
}

func IpInfoCountryAsn(ip net.IP) (interface{}, error) {
	data := cnf.IPINFO_COUNTRY_ASN{}
	return lookupBase(ip, data, cnf.DB_COUNTRY)
}

func IpInfoPrivacy(ip net.IP) (interface{}, error) {
	data := cnf.IPINFO_PRIVACY{}
	return lookupBase(ip, data, cnf.DB_PRIVACY)
}
