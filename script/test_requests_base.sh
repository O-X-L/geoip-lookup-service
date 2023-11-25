#!/bin/bash


echo "${DB_TYPE}: Testing COUNTRY"
curl "http://127.0.0.1:10069/?ip=8.8.8.8&lookup=country"
sleep 1

echo "${DB_TYPE}: Testing ASN"
curl "http://127.0.0.1:10069/?ip=8.8.8.8&lookup=asn"
sleep 1
