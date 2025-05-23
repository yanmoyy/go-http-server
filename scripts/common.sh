#!/bin/bash

# Shared variables
export BASE_URL="http://localhost:8080" # Fixed by adding http://
export HEADERS="Content-Type: application/json"

# Shared function to send curl request
send_post() {
	local test_name="$1"
	local url="$2"
	local data="$3"
	echo "===== $test_name ====="
	echo "URL: $url"
	echo "Data: $data"
	# Capture curl response and HTTP status
	local response
	response=$(curl -s -H "$HEADERS" -d "$data" -X POST "$url" -w "\nStatus Code: %{http_code}")
	echo "Response: $response"
	echo -e "\n-----------------------------------------"
}
export -f send_post
