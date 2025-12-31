#!/usr/bin/env bash
set -e

# Configuration
API="http://localhost:3000"
GUILD_ID="${1:-1}"
EMAIL="charsetadreza@gmail.com"
PASSWORD="reza123"

# Get JWT token
echo "🔑 Getting JWT token..."
TOKEN=$(curl -s -X POST "$API/login" \
  -H "Content-Type: application/json" \
  -d "{\"email\":\"$EMAIL\",\"password\":\"$PASSWORD\"}" | jq -r .token)

if [[ -z "$TOKEN" || "$TOKEN" == "null" ]]; then
  echo "❌ Failed to obtain token"
  exit 1
fi
echo "✅ Token acquired"

# Helper to call API with auth header
curl_api() {
  local method=$1
  local endpoint=$2
  shift 2
  local extra=("$@")
  local headers=("-H" "Content-Type: application/json" "-H" "Authorization: Bearer $TOKEN")
  if [[ "$method" == "POST" || "$method" == "PUT" ]]; then
    headers+=("-d" "$1")
    shift
  fi
  curl -s -X "$method" "$API$endpoint" "${headers[@]}" "$@"
}

# Create channel
echo "📤 Creating channel..."
curl_api POST "/guilds/$GUILD_ID/channels" '{"name":"Test2 Channel"}'
echo

# Fetch channels
echo "📥 Fetching channels..."
curl_api GET "/guilds/$GUILD_ID/channels" | jq
