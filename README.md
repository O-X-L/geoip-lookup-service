[![Lint](https://github.com/O-X-L/geoip-lookup-service/actions/workflows/lint.yml/badge.svg?branch=latest)](https://github.com/O-X-L/geoip-lookup-service/actions/workflows/lint.yml)

# GeoIP Lookup Service

Go-based microservice to perform IP lookups in local GeoIP databases.

It currently only supports databases in **MMDB format**.

If you want to use their extended databases, you might encounter problems. You are welcome to help integrating them correctly.

Feel free to [open a ticket](https://github.com/O-X-L/geoip-lookup-service/issues/new) if you encounter any issues.

----

## GeoIP Provider Support

* **IPInfo**: [Information](https://ipinfo.io/products/free-ip-database), [CC4 License](https://creativecommons.org/licenses/by-sa/4.0/) (*allows for commercial usage - you need to add an attribution*)

    **Attribution**: `<p>IP address data powered by <a href="https://ipinfo.io">IPinfo</a></p>`

* **MaxMind**: [Information](https://dev.maxmind.com/geoip/geolite2-free-geolocation-data), [EULA](https://www.maxmind.com/en/geolite2/eula) (*allows for limited commercial usage - you need to add an attribution*)

    **Attribution**: `This product includes GeoLite2 data created by MaxMind, available from <a href="https://www.maxmind.com">https://www.maxmind.com</a>.`

These two providers were tested.

----

## Integration

* [HAProxy Community using Lua](https://github.com/O-X-L/haproxy-geoip)

   NOTE: HAProxy provides enterprise-grade licensing that has this functionality built-in.

Make sure to read the GeoIP-DB License before integrating it with any service!

----

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

# listen on all external IPs
./geoip_lookup_service -l 0.0.0.0 -p 10069 ...
```

----

## Testing

Basic integration tests are done by using the test-script:

```bash
bash scripts/test.sh
```

Feel free to contribute more test-cases if you found some edge-case issue(s).

----

## Service

Example systemd service:

```text
[Unit]
Description=GeoIP Lookup Service
Documentation=https://github.com/O-X-L/geoip-lookup-service

[Service]
Type=simple
ExecStart=/usr/bin/geoip-lookup -l 127.0.0.1 -p 10069 -t ipinfo -country /etc/geoip/country.mmdb -asn /etc/geoip/asn.mmdb -city /etc/geoip/city.mmdb

# service-user only needs read-access to databases
User=geoip
Group=geoip
Restart=on-failure
RestartSec=5s

StandardOutput=journal
StandardError=journal
SyslogIdentifier=geoip-lookup

[Install]
WantedBy=multi-user.target
```
