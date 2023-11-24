package lookup

import (
	"net"

	"github.com/superstes/geoip-lookup-service/cnf"
)

func IpInfoCountry(ip net.IP, dataStructure interface{}) (interface{}, error) {
	return lookupBase(ip, dataStructure, cnf.DB_COUNTRY)
}

func IpInfoCity(ip net.IP, dataStructure interface{}) (interface{}, error) {
	return lookupBase(ip, dataStructure, cnf.DB_CITY)
}

func IpInfoAsn(ip net.IP, dataStructure interface{}) (interface{}, error) {
	return lookupBase(ip, dataStructure, cnf.DB_ASN)
}

func IpInfoCountryAsn(ip net.IP, dataStructure interface{}) (interface{}, error) {
	return lookupBase(ip, dataStructure, cnf.DB_COUNTRY)
}

func IpInfoPrivacy(ip net.IP, dataStructure interface{}) (interface{}, error) {
	return lookupBase(ip, dataStructure, cnf.DB_PRIVACY)
}
