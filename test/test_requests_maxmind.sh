#!/usr/bin/env bash

source "$(dirname "$0")/test_requests_util.sh"

run_test 'COUNTRY' 'ip=8.8.8.8&lookup=country&filter=*' ''

run_test 'ASN' 'ip=8.8.8.8&lookup=asn&filter=*' '{"autonomous_system_number":15169,"autonomous_system_organization":"GOOGLE"}'

run_test 'COUNTRY with filter' 'ip=8.8.8.8&lookup=country&filter=country.iso_code' 'US'

run_test 'COUNTRY with filter #2' 'ip=8.8.8.8&lookup=country&filter=country.names.en' 'United States'

run_test 'ASN with filter' 'ip=8.8.8.8&lookup=asn&filter=autonomous_system_number' '15169'

run_test 'CITY' 'ip=8.8.8.8&lookup=city' ''

run_test 'CITY with filter' 'ip=8.8.8.8&lookup=city&filter=location' '{"accuracy_radius":1000,"latitude":37.751,"longitude":-97.822,"time_zone":"America/Chicago"}'

run_test 'ASN with filter-shortcut' 'ip=8.8.8.8&lookup=autonomous_system_number' '15169'

run_test 'COUNTRY with filter-shortcut' 'ip=8.8.8.8&lookup=country.names.en' 'United States'
