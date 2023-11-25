#!/bin/bash

echo "${DB_TYPE}: Testing COUNTRY with filter"
curl "http://127.0.0.1:10069/?ip=8.8.8.8&lookup=country&filter=country.iso_code"
sleep 1

echo "${DB_TYPE}: Testing COUNTRY with filter #2"
curl "http://127.0.0.1:10069/?ip=8.8.8.8&lookup=country&filter=country.names.en"
sleep 1

echo "${DB_TYPE}: Testing ASN with filter"
curl "http://127.0.0.1:10069/?ip=8.8.8.8&lookup=asn&filter=autonomous_system_number"
sleep 1

echo "${DB_TYPE}: Testing CITY"
curl "http://127.0.0.1:10069/?ip=8.8.8.8&lookup=city"
sleep 1

echo "${DB_TYPE}: Testing CITY with filter"
curl "http://127.0.0.1:10069/?ip=8.8.8.8&lookup=city&filter=location"
sleep 1
