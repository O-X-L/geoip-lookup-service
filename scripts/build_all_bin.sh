#!/bin/bash

set -euo pipefail

cd "$(dirname "$0")/.."

PATH_BASE="$(pwd)"
PATH_BUILD="${PATH_BASE}/build"
PATH_SRC="${PATH_BASE}/src"

cd "${PATH_BASE}/src"
VERSION="$(cat "./internal/cnf/main.go" | grep VERSION | cut -d '=' -f2 | tr -d ' ')"

mkdir -p "$PATH_BUILD"

rm -f "$PATH_BUILD"/*

APP_NAME="geoip-lookup"

function compile() {
    os="$1" arch="$2"
    cd "$PATH_SRC"
    echo "COMPILING BINARY FOR ${os}-${arch}"
    GOOS="$os" GOARCH="$arch" go build -o "${PATH_BUILD}/${APP_NAME}-${os}-${arch}" cmd/main.go
    GOOS="$os" GOARCH="$arch" CGO_ENABLED=0 go build -o "${PATH_BUILD}/${APP_NAME}-${os}-${arch}-CGO0" cmd/main.go

    cd "$PATH_BUILD"
    if [[ "$os" == "windows" ]]
    then
        zip "./${APP_NAME}-${os}-${arch}.zip" "./${APP_NAME}-${os}-${arch}"
        zip "./${APP_NAME}-${os}-${arch}-CGO0.zip" "./${APP_NAME}-${os}-${arch}-CGO0"
    else
        tar -czf "./${APP_NAME}-${os}-${arch}.tar.gz" "./${APP_NAME}-${os}-${arch}"
        tar -czf "./${APP_NAME}-${os}-${arch}-CGO0.tar.gz" "./${APP_NAME}-${os}-${arch}-CGO0"
    fi
}


compile "linux" "386"
compile "linux" "amd64"
compile "linux" "arm"
compile "linux" "arm64"

# untested
compile "freebsd" "386"
compile "freebsd" "amd64"
compile "freebsd" "arm"

compile "openbsd" "386"
compile "openbsd" "amd64"
compile "openbsd" "arm"

compile "darwin" "amd64"
compile "darwin" "arm64"

compile "windows" "386"
compile "windows" "amd64"

echo "COMMAND TO REMOVE ALL NON-ARCHIVES: find ${PATH_BUILD} -type f ! -name '*.*' -delete"
