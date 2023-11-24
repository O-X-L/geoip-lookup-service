# GeoIP Lookup Service

Go-based microservice to perform IP lookups in local GeoIP databases.

It currently only supports databases in MMDB format.

Providers supported:

* [IPInfo](https://ipinfo.io/account/data-downloads)
* [MaxMind](https://dev.maxmind.com/geoip/geolite2-free-geolocation-data)

As I don't have loads of money to spare - Testing is only done using the [IPInfo Free](https://ipinfo.io/products/free-ip-database) and [MaxMind GeoLite2](https://dev.maxmind.com/geoip/geolite2-free-geolocation-data) databases.

If you want to use their extended databases, you might encounter issues.

## State

This project is still in early development.

You are welcome to test it and report issues you find.


## Usage

The binary starts a simple HTTP webserver.

You can send a query and receive the result as response:


```bash
chmod +x geoip_lookup_service
./geoip_lookup_service -l 127.0.0.1 -p 10069 -t ipinfo -country /etc/geoip/country.mmdb -asn /etc/geoip/asn.mmdb -city /etc/geoip/city.mmdb
# -l = listen address
# -p = listen port
# -t = database type (ipinfo/maxmind)
# -country = path to country-database
# -city = path to city-database
# -asn = path to asn-database

curl http://127.0.0.1:10069/?lookup=country&ip=1.1.1.1
> ...

curl http://127.0.0.1:10069/?lookup=country&ip=1.1.1.1&filter=xxx
> ...

curl http://127.0.0.1:10069/?lookup=asn&ip=1.1.1.1
> ...

curl http://127.0.0.1:10069/?lookup=city&ip=1.1.1.1
> ...
```
