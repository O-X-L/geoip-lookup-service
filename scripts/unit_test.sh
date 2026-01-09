#!/usr/bin/env bash

set -euo pipefail

cd "$(dirname "$0")/../src"

export MODE_TEST=1

go run gotest.tools/gotestsum@latest --format pkgname ./...
