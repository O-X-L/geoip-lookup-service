#!/bin/bash

echo "${DB_TYPE}: Testing COUNTRY with filter"
curl "http://127.0.0.1:10069/?ip=8.8.8.8&lookup=country&filter=country"
sleep 1

echo "${DB_TYPE}: Testing ASN with filter"
curl "http://127.0.0.1:10069/?ip=8.8.8.8&lookup=asn&filter=asn"
sleep 1

