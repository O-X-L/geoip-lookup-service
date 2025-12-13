#!/bin/bash

set -eo pipefail

PORT=10069

DB_TYPE='ipinfo'
if [ -n "$1" ]
then
  DB_TYPE="$1"
fi

PATH_DB="${HOME}/Downloads"
if [ -n "$2" ]
then
  PATH_DB="$2"
fi

set -u

binary="/tmp/geoip_lookup_$(date +"%s")"

cd "$(dirname "$0")/../src"
go build -o "$binary" cmd/main.go
chmod +x "$binary"

echo "RUNNING GeoIP Lookup: ${DB_TYPE}"

if [[ "$DB_TYPE" == "ipinfo" ]]
then
  DB_LITE="${PATH_DB}/ipinfo_lite.mmdb"
  if ! [ -f "$DB_LITE" ]
  then
    echo "ERROR: MMDB-File does not exist => ${DB_LITE}"
    exit 1
  fi
  "$binary" -t "$DB_TYPE" -p "$PORT" -lite "$DB_LITE"

elif [[ "$DB_TYPE" == "maxmind" ]]
then
  DB_ASN="${PATH_DB}/maxmind_asn.mmdb"
  DB_COUNTRY="${PATH_DB}/maxmind_country.mmdb"
  DB_CITY="${PATH_DB}/maxmind_city.mmdb"
  if ! [ -f "$DB_ASN" ] || ! [ -f "$DB_COUNTRY" ] || ! [ -f "$DB_CITY" ]
  then
    echo "ERROR: One or more MMDB-Files do not exist => ${DB_ASN}, ${DB_COUNTRY}, ${DB_CITY}"
    exit 1
  fi
  "$binary" -t "$DB_TYPE" -p "$PORT" -country "$DB_COUNTRY" -asn "$DB_ASN" -city "$DB_CITY"

else
  DB_ASN="${PATH_DB}/oxl_asn.mmdb"
  if ! [ -f "$DB_ASN" ]
  then
    echo "ERROR: MMDB-File does not exist => ${DB_ASN}"
    exit 1
  fi
  "$binary" -t "$DB_TYPE" -p "$PORT" -asn "$DB_ASN"
fi
