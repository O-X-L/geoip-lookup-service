package lookup

import (
	"net"

	"github.com/superstes/geoip-lookup-service/cnf"
)

func MaxMindCountry(ip net.IP, dataStructure interface{}) (interface{}, error) {
	return lookupBase(ip, dataStructure, cnf.DB_COUNTRY)
}

func MaxMindCity(ip net.IP, dataStructure interface{}) (interface{}, error) {
	return lookupBase(ip, dataStructure, cnf.DB_CITY)
}

func MaxMindAsn(ip net.IP, dataStructure interface{}) (interface{}, error) {
	return lookupBase(ip, dataStructure, cnf.DB_ASN)
}

func MaxMindPrivacy(ip net.IP, dataStructure interface{}) (interface{}, error) {
	return lookupBase(ip, dataStructure, cnf.DB_PRIVACY)
}
