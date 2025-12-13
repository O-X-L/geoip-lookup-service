package lookup

import (
	"net"

	"git.oxl.at/geoip-lookup-service/internal/cnf"
	"github.com/oschwald/maxminddb-golang"
)

var FUNC_MAPPING = map[uint]map[string]func(net.IP) (interface{}, error){
	cnf.DB_TYPE_NONE: {},
	cnf.DB_TYPE_IPINFO: {
		"country_asn": IpInfoCountryAsn,
		"country":     IpInfoCountry,
		"city":        IpInfoCity,
		"asn":         IpInfoAsn,
		"privacy":     IpInfoPrivacy,
	},
	cnf.DB_TYPE_MAXMIND: {
		"country_asn": nil,
		"country":     MaxMindCountry,
		"city":        MaxMindCity,
		"asn":         MaxMindAsn,
		"privacy":     MaxMindPrivacy,
	},
	cnf.DB_TYPE_OXL: {
		"asn": OXLAsn,
	},
}

func lookupBase(ip net.IP, dataStructure interface{}, dbFile string) (interface{}, error) {
	db, err := maxminddb.Open(dbFile)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.Lookup(ip, &dataStructure)
	if err != nil {
		return nil, err
	}
	return dataStructure, nil
}
