#!/usr/bin/env bash

source "$(dirname "$0")/test_requests_util.sh"

run_test 'ASN' 'ip=8.8.8.8&lookup=asn&filter=*' '{"asn":15169,"contacts":{},"info":{"name":"Google LLC","name_long":"","website":"https://about.google/intl/en/"},"ipv4_count":1946112,"ipv6_count":714059288910298409815251091456,"organization":{"city":"Mountain View","country":"US","latitude":0,"longitude":0,"name":"Google LLC","name_long":"","state":"California","website":"http://www.google.com"}}'

run_test 'ASN with filter' 'ip=1.1.1.1&lookup=asn&filter=asn' '13335'

run_test 'ASN with filter for org-name' 'ip=1.1.1.1&lookup=asn&filter=organization.name' 'Cloudflare, Inc.'

run_test 'ASN with filter for org-country' 'ip=1.1.1.1&lookup=asn&filter=organization.country' 'US'

run_test 'ASN with filter for abuse-contact' 'ip=1.1.1.1&lookup=asn&filter=contacts.abuse.email' 'abuse@cloudflare.com'

run_test 'ASN with filter-shortcut' 'ip=1.1.1.1&lookup=asn' '13335'

run_test 'ASN with filter-shortcut #2' 'ip=1.1.1.1&lookup=organization.name' 'Cloudflare, Inc.'
