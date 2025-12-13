#!/usr/bin/env bash

source "$(dirname "$0")/test_requests_util.sh"

run_test 'COUNTRY' 'ip=8.8.8.8&lookup=country&filter=*' '{"as_domain":"google.com","as_name":"Google LLC","asn":"AS15169","continent":"North America","continent_code":"NA","country":"United States","country_code":"US"}'

run_test 'ASN' 'ip=8.8.8.8&lookup=asn&filter=*' '{"as_domain":"google.com","as_name":"Google LLC","asn":"AS15169","continent":"North America","continent_code":"NA","country":"United States","country_code":"US"}'

run_test 'COUNTRY with filter' 'ip=8.8.8.8&lookup=country&filter=country' 'United States'

run_test 'ASN with filter' 'ip=8.8.8.8&lookup=asn&filter=asn' 'AS15169'

run_test 'COUNTRY with filter-shortcut' 'ip=8.8.8.8&lookup=country' 'United States'
