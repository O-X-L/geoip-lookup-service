package cnf

import "net"

const DB_TYPE_OXL uint8 = 3

type contact struct {
	Name  string `maxminddb:"name"`
	Email string `maxminddb:"email"`
	Phone string `maxminddb:"phone"`
	URL   string `maxminddb:"url"`
}

type socialMedia struct {
	Service    string `maxminddb:"service"`
	Identifier string `maxminddb:"identifier"`
}

// OXL GEOIP-ASN schema: https://github.com/O-X-L/geoip-asn/blob/latest/schema | https://github.com/O-X-L/geoip-asn/blob/latest/example
var OXL_ASN struct {
	ASN          int `maxminddb:"asn"`
	IPv4Count    int `maxminddb:"ipv4_count"`
	IPv6Count    int `maxminddb:"ipv6_count"`
	Organization struct {
		Status      string        `maxminddb:"status"`
		Address1    string        `maxminddb:"address1"`
		Address2    string        `maxminddb:"address2"`
		City        string        `maxminddb:"city"`
		State       string        `maxminddb:"state"`
		ZIPCode     string        `maxminddb:"zipcode"`
		Country     string        `maxminddb:"country"`
		Suite       string        `maxminddb:"suite"`
		Floor       string        `maxminddb:"floor"`
		Latitude    float64       `maxminddb:"latitude"`
		Longitude   float64       `maxminddb:"longitude"`
		Name        string        `maxminddb:"name"`
		AKA         string        `maxminddb:"aka"`
		NameLong    string        `maxminddb:"name_long"`
		Website     string        `maxminddb:"website"`
		SocialMedia []socialMedia `maxminddb:"social_media"`
		Notes       string        `maxminddb:"notes"`
	} `maxminddb:"organization"`
	Info struct {
		Status                   string        `maxminddb:"status"`
		Name                     string        `maxminddb:"name"`
		AKA                      string        `maxminddb:"aka"`
		NameLong                 string        `maxminddb:"name_long"`
		IrrAsSet                 string        `maxminddb:"irr_as_set"`
		Website                  string        `maxminddb:"website"`
		SocialMedia              []socialMedia `maxminddb:"rir_status_updated"`
		LookingGlass             string        `maxminddb:"looking_glass"`
		RouteServer              string        `maxminddb:"route_server"`
		Notes                    string        `maxminddb:"notes"`
		InfoTraffic              string        `maxminddb:"info_traffic"`
		InfoRatio                string        `maxminddb:"info_ratio"`
		InfoScope                string        `maxminddb:"info_scope"`
		InfoTypes                []string      `maxminddb:"info_types"`
		InfoPrefixes4            int           `maxminddb:"info_prefixes4"`
		InfoPrefixes6            int           `maxminddb:"info_prefixes6"`
		InfoUnicast              bool          `maxminddb:"info_unicast"`
		InfoMulticast            bool          `maxminddb:"info_multicast"`
		InfoIPv6                 bool          `maxminddb:"info_ipv6"`
		InfoNeverViaRouteServers bool          `maxminddb:"info_never_via_route_servers"`
		PolicyURL                string        `maxminddb:"policy_url"`
		PolicyGeneral            string        `maxminddb:"policy_general"`
		PolicyLocations          string        `maxminddb:"policy_locations"`
		PolicyRatio              bool          `maxminddb:"policy_ratio"`
		PolicyContracts          string        `maxminddb:"policy_contracts"`
		StatusDashboard          string        `maxminddb:"status_dashboard"`
		RirStatus                string        `maxminddb:"status_dashboard"`
		RirStatusUpdated         string        `maxminddb:"rir_status"`
	} `maxminddb:"info"`
	Contacts struct {
		NOC    contact `maxminddb:"noc"`
		Policy contact `maxminddb:"policy"`
		Abuse  contact `maxminddb:"abuse"`
	} `maxminddb:"contacts"`

	Network       net.IP `maxminddb:"network"`
	Country       string `maxminddb:"country_code"`
	CountryName   string `maxminddb:"country"`
	Continent     string `maxminddb:"continent_code"`
	ContinentName string `maxminddb:"continent"`
	Name          string `maxminddb:"as_name"`
	Domain        string `maxminddb:"as_domain"`
}
