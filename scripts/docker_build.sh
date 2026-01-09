#!/usr/bin/env bash

if [ -z "$1" ]
then
  echo "USAGE:"
  echo " 1 > Version"
  exit 1
fi

set -euo pipefail

VERSION="$1"

cd "$(dirname "$0")/../docker/default"

docker build -f Dockerfile -t "oxlorg/geoip-lookup:${VERSION}" --network=host --no-cache ../..
docker build -f Dockerfile -t "oxlorg/geoip-lookup:latest" --network=host ../..
