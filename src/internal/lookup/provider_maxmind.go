package lookup

import (
	"net"

	"git.oxl.at/geoip-lookup-service/internal/cnf"
)

func MaxMindCountry(ip net.IP) (interface{}, error) {
	data := cnf.MAXMIND_COUNTRY{}
	return lookupBase(ip, data, cnf.DB_COUNTRY)
}

func MaxMindCity(ip net.IP) (interface{}, error) {
	data := cnf.MAXMIND_CITY{}
	return lookupBase(ip, data, cnf.DB_CITY)
}

func MaxMindAsn(ip net.IP) (interface{}, error) {
	data := cnf.MAXMIND_ASN{}
	return lookupBase(ip, data, cnf.DB_ASN)
}

func MaxMindPrivacy(ip net.IP) (interface{}, error) {
	data := cnf.MAXMIND_PRIVACY{}
	return lookupBase(ip, data, cnf.DB_PRIVACY)
}
