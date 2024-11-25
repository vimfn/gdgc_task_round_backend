#!/bin/bash

# Base URL for the API (replace with your actual API endpoint)
BASE_URL="http://localhost:8080"

# Create a new listing
echo "1. Creating a new listing:"
CREATE_RESPONSE=$(curl -s -X POST "$BASE_URL/listing" \
  -H "Content-Type: application/json" \
  -d '{"title":"Test Item","description":"A test item for API","seller":"TestSeller","rating":4.5}' \
  | jq '.')
echo "$CREATE_RESPONSE"

# Extract the ID from the created listing
ITEM_ID=$(echo "$CREATE_RESPONSE" | jq -r '.data.id')

# Get all listings
echo -e "\n2. Getting all listings:"
curl -s "$BASE_URL/listing" | jq '.'

# Get specific listing by ID
echo -e "\n3. Getting specific listing:"
curl -s "$BASE_URL/listing/$ITEM_ID" | jq '.'

# Update the listing
echo -e "\n4. Updating the listing:"
UPDATE_RESPONSE=$(curl -s -X PUT "$BASE_URL/listing/$ITEM_ID" \
  -H "Content-Type: application/json" \
  -d '{"title":"Updated Test Item","description":"An updated test item"}' \
  | jq '.')
echo "$UPDATE_RESPONSE"

# Delete the listing
echo -e "\n5. Deleting the listing:"
curl -s -X DELETE "$BASE_URL/listing/$ITEM_ID" -I
