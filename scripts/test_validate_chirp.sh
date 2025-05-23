#!/bin/bash

# Define common variables
ENDPOINT="localhost:8080/api/validate_chirp"
HEADERS="Content-Type: application/json"

# Function to send curl request
send_chirp() {
	local test_name="$1"
	local chirp_data="$2"
	echo "Running $test_name..."
	curl -H "$HEADERS" -d "$chirp_data" -X POST "$ENDPOINT"
	echo -e "\n-----------------------------------------"
}

# @lang json
SHORT='{
  "body": "I had something interesting for breakfast"
}'
send_chirp "Test 1: Sending a short chirp" "$SHORT"

# @lang json
EXTRA_FIELD='{
  "body": "I hear Mastodon is better than Chirpy. sharbert I need to migrate",
  "extra": "this should be ignored"
}'
send_chirp "Test 3: Sending a extra field" "$EXTRA_FIELD"

# @lang json
BAD_WORD='{
  "body": "I really need a kerfuffle to go to bed sooner, Fornax !"
}'
send_chirp "Test 4: Sending a bad word" "$BAD_WORD"

# @lang json
LONG_CHIRP='{
  "body": "lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
}'
send_chirp "Test 2: Sending a long chirp" "$LONG_CHIRP"
