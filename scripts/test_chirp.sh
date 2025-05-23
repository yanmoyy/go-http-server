#!/bin/bash

SCRIPT_DIR=$(dirname "$0")
source "$SCRIPT_DIR/common.sh"

CHIRP_URL="$BASE_URL/api/validate_chirp"

# @lang json
SHORT='{
  "body": "I had something interesting for breakfast"
}'
send_post "Test 1: Sending a short chirp" "$CHIRP_URL" "$SHORT"

# @lang json
EXTRA_FIELD='{
  "body": "I hear Mastodon is better than Chirpy. sharbert I need to migrate",
  "extra": "this should be ignored"
}'
send_post "Test 3: Sending a extra field" "$CHIRP_URL" "$EXTRA_FIELD"
# @lang json
BAD_WORD='{
  "body": "I really need a kerfuffle to go to bed sooner, Fornax !"
}'
send_post "Test 4: Sending a bad word" "$CHIRP_URL" "$BAD_WORD"
# @lang json
LONG_CHIRP='{
  "body": "lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
}'
send_post "Test 2: Sending a long chirp" "$CHIRP_URL" "$LONG_CHIRP"
