package cnf

var LOOKUP_FILTER_SHORTCUTS = map[uint]map[string]string{
	DB_TYPE_NONE: {},
	DB_TYPE_IPINFO: {
		"country_code":   "country",
		"country":        "country",
		"asn":            "asn",
		"as_name":        "asn",
		"continent":      "country",
		"continent_name": "country",
	},
	DB_TYPE_MAXMIND: {
		"country.iso_code":               "country",
		"country.names.en":               "country",
		"autonomous_system_number":       "asn",
		"autonomous_system_organization": "asn",
		"isp":                            "asn",
		"organization":                   "asn",
		"continent.code":                 "country",
		"continent.names.en":             "country",
	},
	DB_TYPE_OXL: {
		"asn":                  "asn",
		"organization.name":    "asn",
		"organization.country": "asn",
	},
}
