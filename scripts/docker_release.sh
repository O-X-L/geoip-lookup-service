#!/usr/bin/env bash

if [ -z "$1" ]
then
  echo "USAGE:"
  echo " 1 > Version"
  exit 1
fi

set -euo pipefail

VERSION="$1"

docker push "oxlorg/geoip-lookup:${VERSION}"
docker push "oxlorg/geoip-lookup:latest"
