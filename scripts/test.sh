#!/bin/bash

if [ -z "$1" ]
then
  PATH_DB="${HOME}/Downloads"
else
  PATH_DB="$1"
fi

set -euo pipefail

cd "$(dirname "$0")"

function kill_ps() {
  pkill -f "/tmp/geoip_lookup_*" > /dev/null || true
  sleep 1
}

kill_ps

echo ''
echo "### TESTING IPINFO DATABASES ###"

./build_run.sh "ipinfo" "$PATH_DB" > /dev/null &
sleep 1
export DB_TYPE="IPINFO"
./test_requests_ipinfo.sh

kill_ps

echo ''
echo "### TESTING MAXMIND DATABASES ###"

./build_run.sh "maxmind" "$PATH_DB" > /dev/null &
sleep 1
export DB_TYPE="MAXMIND"
./test_requests_maxmind.sh

kill_ps

echo ''
echo "### TESTING OXL DATABASES ###"

./build_run.sh "oxl" "$PATH_DB" > /dev/null &
sleep 1
export DB_TYPE="OXL"
./test_requests_oxl.sh

kill_ps

echo ''
echo 'FINISHED'
echo ''
