#!/bin/bash

set -eo pipefail

PORT=10069

DB_TYPE='ipinfo'
if [ -n "$1" ]
then
  DB_TYPE="$1"
fi

DB_DIR='/etc/geoip'
if [ -n "$2" ]
then
  DB_DIR="$2"
fi

set -u

DB_CO="${DB_DIR}/${DB_TYPE}_country.mmdb"
DB_AS="${DB_DIR}/${DB_TYPE}_asn.mmdb"
DB_CI="${DB_DIR}/${DB_TYPE}_city.mmdb"

if ! [ -f "$DB_CO" ] || ! [ -f "$DB_AS" ] || ([[ "$DB_TYPE" == "maxmind" ]] && ! [ -f "$DB_CI" ])
then
  echo "ERROR: Required databases missing (${DB_DIR}/${DB_TYPE}_*)"
  exit 1
fi



binary="/tmp/geoip_lookup_$(date +"%s")"

cd "$(dirname "$0")/../main"
go build -o "$binary"
chmod +x "$binary"

echo "RUNNING GeoIP Lookup: ${DB_DIR}/${DB_TYPE}_*"
"$binary" -t "$DB_TYPE" -p "$PORT" -country "$DB_CO" -asn "$DB_AS" -city "$DB_CI"
