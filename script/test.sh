#!/bin/bash

set -euo pipefail

cd "$(dirname "$0")"

echo ''
echo "### TESTING IPINFO DATBASES ###"

./build_run.sh "ipinfo"  > /dev/null &
sleep 1
export DB_TYPE="IPINFO"
./test_requests_base.sh
./test_requests_ipinfo.sh

pkill -f "/tmp/geoip_lookup_*" > /dev/null
sleep 1

echo ''
echo "### TESTING MAXMIND DATBASES ###"

./build_run.sh "maxmind" > /dev/null &
sleep 1
export DB_TYPE="MAXMIND"
./test_requests_base.sh
./test_requests_maxmind.sh

pkill -f "/tmp/geoip_lookup_*" > /dev/null
sleep 1

echo ''
echo 'FINISHED'
echo ''
