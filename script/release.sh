#!/bin/bash

set -euo pipefail

cd "$(dirname "$0")/../main"

VERSION="$(cat "../cnf/main.go" | grep VERSION | cut -d '=' -f2 | tr -d ' ')"

mkdir -p "../build"

rm ../build/*

function compile() {
    os="$1" arch="$2"
    echo "COMPILING BINARY FOR ${os}-${arch}"
    GOOS="$os" GOARCH="$arch" go build -o "../build/geoip-lookup-${os}-${arch}"
    GOOS="$os" GOARCH="$arch" CGO_ENABLED=0 go build -o "../build/geoip-lookup-${os}-${arch}-CGO0"
}

compile "linux" "386"
compile "linux" "amd64"
compile "linux" "arm"
compile "linux" "arm64"

compile "freebsd" "386"
compile "freebsd" "amd64"
compile "freebsd" "arm"

compile "openbsd" "386"
compile "openbsd" "amd64"
compile "openbsd" "arm"

compile "darwin" "amd64"
compile "darwin" "arm64"

# untested
compile "windows" "386"
compile "windows" "amd64"
