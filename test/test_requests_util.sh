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
    echo -e "\033[0;31m > ERROR: Output not as expected => ${expect} \033[0m"
  else
    echo -e "\033[0;32m > OK \033[0m"
  fi
  echo ''

  sleep 1
}
