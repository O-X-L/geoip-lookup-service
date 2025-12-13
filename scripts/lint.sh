#!/usr/bin/env bash

if [ -z "$CI" ]
then
  CI="0"
fi

if [ -z "$1" ]
then
  set -euo pipefail
else
  set -uo pipefail
fi

cd "$(dirname "$0")/.."

BASE_DIR="$(pwd)"

function lint() {
  echo -e '\033[0;33m'
  echo 'WARNINGS:'
  echo ''
  warnings="$(golangci-lint --config="${BASE_DIR}/.golangci_warn.yml" run || true)"
  if [[ "$CI" == "0" ]]
  then
    echo "$warnings"
  else
    # should be shown as warning in github-CI action-overview
    echo "::warning $warnings"
  fi
  echo -e '\033[0m'
  echo -e '\033[0;31m'
  echo ''
  echo 'ERRORS:'
  echo ''
  golangci-lint --config="${BASE_DIR}/.golangci_fail.yml" run
  echo -e '\033[0m'
}

cd ./src/
lint
