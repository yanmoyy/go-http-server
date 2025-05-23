#!/bin/bash

ORIGINAL_DIR=$(pwd)
cd "$(dirname "$0")" || exit 1
source ./common.sh
cd "$ORIGINAL_DIR" || exit 1

# Test cases for user creation
send_post "Test 0: Reset", "$URL_RESET"

# @lang json
user='{
  "email": "john@example.com"
}'
send_post "Test 1: Creating a valid user" "$URL_USER" "$user"

# @lang json
user='{
  "email": "dackjorsey@example.co"
}'
send_post "Test 2: Creating a valid user" "$URL_USER" "$user"
