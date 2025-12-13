package lookup

import (
	"net"

	"git.oxl.at/geoip-lookup-service/internal/cnf"
)

func OXLAsn(ip net.IP) (interface{}, error) {
	data := cnf.OXL_ASN{}
	return lookupBase(ip, data, cnf.DB_ASN)
}
