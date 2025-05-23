#!/bin/bash

SCRIPT_DIR=$(dirname "$0")
source "$SCRIPT_DIR/common.sh"

# Set endpoint for user creation
URL_USER="$BASE_URL/api/users"
URL_RESET="$BASE_URL/admin/reset"

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
