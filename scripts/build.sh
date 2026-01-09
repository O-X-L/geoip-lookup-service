#!/bin/bash

set -euo pipefail

cd "$(dirname "$0")/.."

BASE_DIR="$(pwd)"
mkdir -p "${BASE_DIR}/build"
cd ./src/
go build -o "${BASE_DIR}/build/geoip-lookup" cmd/main.go
