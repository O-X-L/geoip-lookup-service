package lookup

import (
	"net"

	"git.oxl.at/geoip-lookup-service/internal/cnf"
)

func MaxMindCountry(ip net.IP) (interface{}, error) {
	return lookupBase(ip, cnf.MAXMIND_COUNTRY, cnf.DB_COUNTRY)
}

func MaxMindCity(ip net.IP) (interface{}, error) {
	return lookupBase(ip, cnf.MAXMIND_CITY, cnf.DB_CITY)
}

func MaxMindAsn(ip net.IP) (interface{}, error) {
	return lookupBase(ip, cnf.MAXMIND_ASN, cnf.DB_ASN)
}

func MaxMindPrivacy(ip net.IP) (interface{}, error) {
	return lookupBase(ip, cnf.MAXMIND_PRIVACY, cnf.DB_PRIVACY)
}
