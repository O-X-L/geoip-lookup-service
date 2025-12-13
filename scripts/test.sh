#!/bin/bash

echo "INFO: For MMDB-download-links see => https://github.com/O-X-L/geoip-lookup-service/blob/latest/.github/workflows/test.yml#L40"

if [ -z "$1" ]
then
  PATH_DB="${HOME}/Downloads"
else
  PATH_DB="$1"
fi

set -euo pipefail

cd "$(dirname "$0")/.."
PATH_BASE="$(pwd)"

function kill_ps() {
  pkill -f "/tmp/geoip_lookup_*" > /dev/null || true
  sleep 1
}

kill_ps

echo ''
echo "### TESTING IPINFO DATABASES ###"

bash "${PATH_BASE}/scripts/build_run.sh" "ipinfo" "$PATH_DB" > /dev/null &
sleep 1
export DB_TYPE="IPINFO"
bash "${PATH_BASE}/test/test_requests_ipinfo.sh"

kill_ps

echo ''
echo "### TESTING MAXMIND DATABASES ###"

bash "${PATH_BASE}/scripts/build_run.sh" "maxmind" "$PATH_DB" > /dev/null &
sleep 1
export DB_TYPE="MAXMIND"
bash "${PATH_BASE}/test/test_requests_maxmind.sh"

kill_ps

echo ''
echo "### TESTING OXL DATABASES ###"

bash "${PATH_BASE}/scripts/build_run.sh" "oxl" "$PATH_DB" > /dev/null &
sleep 1
export DB_TYPE="OXL"
bash "${PATH_BASE}/test/test_requests_oxl.sh"

kill_ps

echo ''
echo 'FINISHED'
echo ''
