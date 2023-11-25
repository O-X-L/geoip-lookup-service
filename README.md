[![Lint](https://github.com/superstes/geoip-lookup-service/actions/workflows/lint.yml/badge.svg?branch=latest)](https://github.com/superstes/geoip-lookup-service/actions/workflows/lint.yml)

# GeoIP Lookup Service

Go-based microservice to perform IP lookups in local GeoIP databases.

It currently only supports databases in MMDB format.

Feel free to [open a ticket](https://github.com/superstes/geoip-lookup-service/issues/new) if you encounter any issues.

Providers supported:

* [IPInfo](https://ipinfo.io/account/data-downloads)
* [MaxMind](https://dev.maxmind.com/geoip/geolite2-free-geolocation-data)

As I don't have loads of money to spare - Testing is only done using the [IPInfo Free](https://ipinfo.io/products/free-ip-database) and [MaxMind GeoLite2](https://dev.maxmind.com/geoip/geolite2-free-geolocation-data) databases.

If you want to use their extended databases, you might encounter problems. You are welcome to help integrating them correctly.

## Usage

The binary starts a simple HTTP webserver.

You can send a query and receive the result as response:


```bash
chmod +x geoip_lookup_service
./geoip_lookup_service -l 127.0.0.1 -p 10069 -t ipinfo -country /etc/geoip/country.mmdb -asn /etc/geoip/asn.mmdb -city /etc/geoip/city.mmdb
# -l = listen address (default=127.0.0.1)
# -p = listen port (default=10000)
# -plain = response in plain text format (default=false)
# -t = database type (ipinfo/maxmind) (default=ipinfo)
# -country = path to country-database (default=/etc/geoip/country.mmdb)
# -city = path to city-database (default=/etc/geoip/city.mmdb)
# -asn = path to asn-database (default=/etc/geoip/asn.mmdb)

curl "http://127.0.0.1:10069/?lookup=country&ip=1.1.1.1"
> {"continent":"NA","continent_name":"North America","country":"US","country_name":"United States"}

curl "http://127.0.0.1:10069/?lookup=country&ip=1.1.1.1&filter=country"
> "US"

curl "http://127.0.0.1:10069/?lookup=asn&ip=1.1.1.1"
> {"asn":"AS13335","domain":"cloudflare.com","name":"Cloudflare, Inc."}

# use the 'plain' flag to get single attributes without JSON formatting
./geoip_lookup_service -plain ...
curl "http://127.0.0.1:10069/?lookup=country&ip=1.1.1.1&filter=country_name"
> United States

# use other DB-type
./geoip_lookup_service -t maxmind ...

curl "http://127.0.0.1:10069/?ip=1.1.1.1&lookup=asn"
> {"autonomous_system_number":13335,"autonomous_system_organization":"CLOUDFLARENET"}

curl "http://127.0.0.1:10069/?ip=1.1.1.1&lookup=country"
> {"registered_country":{"geoname_id":2077456,"iso_code":"AU","names":{"de":"Australien","en":"Australia","es":"Australia","fr":"Australie","ja":"オーストラリア","pt-BR":"Austrália","ru":"Австралия","zh-CN":"澳大利亚"}}}

# filters can also be used to get deeper attributes
curl "http://127.0.0.1:10069/?ip=8.8.8.8&lookup=country"
> {"continent":{"code":"NA","geoname_id":6255149,"names":{"de":"Nordamerika","en":"North America","es":"Norteamérica","fr":"Amérique du Nord","ja":"北アメリカ","pt-BR":"América do Norte","ru":"Северная Америка","zh-CN":"北美洲"}},"country":{"geoname_id":6252001,"iso_code":"US","names":{"de":"Vereinigte Staaten","en":"United States","es":"Estados Unidos","fr":"États Unis","ja":"アメリカ","pt-BR":"EUA","ru":"США","zh-CN":"美国"}},"registered_country":{"geoname_id":6252001,"iso_code":"US","names":{"de":"Vereinigte Staaten","en":"United States","es":"Estados Unidos","fr":"États Unis","ja":"アメリカ","pt-BR":"EUA","ru":"США","zh-CN":"美国"}}}

curl "http://127.0.0.1:10069/?ip=8.8.8.8&lookup=country&filter=country.iso_code"
> "US"

curl "http://127.0.0.1:10069/?ip=8.8.8.8&lookup=country&filter=country.names.en"
> "United States"

curl "http://127.0.0.1:10069/?ip=8.8.8.8&lookup=city&filter=location"
> {"accuracy_radius":1000,"latitude":37.751,"longitude":-97.822,"time_zone":"America/Chicago"}
```

----

## Testing

Basic integration tests are done by using the test-script:

```bash
bash scripts/test.sh
```

Feel free to contribute more test-cases if you found some edge-case issue(s).
