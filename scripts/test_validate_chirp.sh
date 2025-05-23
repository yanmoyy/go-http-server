#!/bin/bash

# Define common variables
ENDPOINT="localhost:8080/api/validate_chirp"
HEADERS="Content-Type: application/json"

# asdfdsf
SHORT3='{"body": "I had something interesting for breakfast"}'
# Test 1: Short chirp

# @lang json
SHORT='
{
  "body": "I had something interesting for breakfast",
  "basdf": "sdfdfsf"
}'
echo "Running Test 1: Sending a short chirp..."
curl -H "$HEADERS" \
	-d "$SHORT" \
	-X POST \
	"$ENDPOINT"
echo "-----------------------------------------"

# Test 2: Long chirp (lorem ipsum)
echo "Running Test 2: Sending a long chirp..."
# Use a here-document for better JSON readability
# Tip: this is JSON

# @lang json
LONG_CHIRP='
{
  "body": "lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
}
'
curl -H "$HEADERS" \
	-d "$LONG_CHIRP" \
	-X POST \
	"$ENDPOINT"
echo "-----------------------------------------"
