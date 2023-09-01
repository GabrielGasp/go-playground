#!/bin/bash

BASE_URL="https://onabet.com/api/gs/v2/game"
MAX_SLICE=218
OUTPUT_FILE="icon_urls.txt"

function make_request {
  local url="$BASE_URL?currentSlice=$CURRENT_SLICE&group=CASINO&name="
  local response=$(curl -s "$url" \
    -H 'authority: onabet.com' \
    -H 'accept: application/json' \
    -H 'accept-language: pt-BR,pt;q=0.6' \
    -H 'content-type: application/json' \
    -H 'cookie: cf_clearance=UoicelNrnPyK1Y1kUbfU8WxVdY0OK3oq637FUBxSP90-1693590407-0-1-f9361fc2.6e9d5a88.2ca69a57-0.2.1693590407; __cf_bm=83tHDANXFg4u0GU7iGsVZNBVtKklTq885Z5LldUYFH8-1693593567-0-AeTbKfTwmN9bO1sQ0hNL1boyKnSPv9s5AjbadtRNMxz1Lfp3WJZkKAO9HbWR+4pduMmbVawxO4T3F3xGE2/b0jU=' \
    -H 'device: desktop' \
    -H 'referer: https://onabet.com/casino' \
    -H 'sec-ch-ua: "Chromium";v="116", "Not)A;Brand";v="24", "Brave";v="116"' \
    -H 'sec-ch-ua-mobile: ?0' \
    -H 'sec-ch-ua-platform: "macOS"' \
    -H 'sec-fetch-dest: empty' \
    -H 'sec-fetch-mode: cors' \
    -H 'sec-fetch-site: same-origin' \
    -H 'sec-gpc: 1' \
    -H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36' \
    -H 'x-locale: BR_PT' \
    -H 'x-project-id: 168' \
    --compressed)
  local icon_urls=$(echo "$response" | jq -r '.games[].iconUrl')

  if [ -n "$icon_urls" ]; then
    echo "$icon_urls" >> "$OUTPUT_FILE"
  fi
}

for ((CURRENT_SLICE = 0; CURRENT_SLICE <= MAX_SLICE; CURRENT_SLICE++)); do
  make_request

  sleep 2
done
