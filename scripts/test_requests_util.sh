#!/usr/bin/env bash

function run_test() {
  comment="$1"
  query="$2"
  expect="$3"

  echo "${DB_TYPE}: Testing ${comment}"
  result="$(curl "http://127.0.0.1:10069/?${query}" 2>/dev/null)"

  echo "RESULT: ${result}"
  if [[ "$expect" != "" ]] && [[ "$result" != "$expect" ]]
  then
    echo "ERROR: Output not as expected => ${expect}"
  fi
  echo ''

  sleep 1
}
