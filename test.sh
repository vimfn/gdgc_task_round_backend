#!/bin/bash

# included jwt auth

BASE_URL="http://localhost:8080"

# User Registration
echo "1. Registering a new user:"
REGISTER_RESPONSE=$(curl -s -X POST "$BASE_URL/register" \
  -H "Content-Type: application/json" \
  -d '{"email":"testuser@example.com","password":"testpassword"}')
echo "$REGISTER_RESPONSE"

# User Login to get JWT token
echo -e "\n2. Logging in and getting JWT token:"
LOGIN_RESPONSE=$(curl -s -X POST "$BASE_URL/login" \
  -H "Content-Type: application/json" \
  -d '{"email":"testuser@example.com","password":"testpassword"}')
TOKEN=$(echo "$LOGIN_RESPONSE" | jq -r '.token')
echo "Token: $TOKEN"

# Create a new listing with JWT auth
echo -e "\n3. Creating a new listing:"
CREATE_RESPONSE=$(curl -s -X POST "$BASE_URL/listing" \
  -H "Content-Type: application/json" \
  -H "Authorization: $TOKEN" \
  -d '{"title":"Test Item","description":"A test item for API","seller":"TestSeller","rating":4.5}' \
  | jq '.')
echo "$CREATE_RESPONSE"

# Extract the ID from the created listing
ITEM_ID=$(echo "$CREATE_RESPONSE" | jq -r '.data.id')

# Get all listings (public route)
echo -e "\n4. Getting all listings:"
curl -s "$BASE_URL/listing" | jq '.'

# Get specific listing by ID (public route)
echo -e "\n5. Getting specific listing:"
curl -s "$BASE_URL/listing/$ITEM_ID" | jq '.'

# Update the listing with JWT auth
echo -e "\n6. Updating the listing:"
UPDATE_RESPONSE=$(curl -s -X PUT "$BASE_URL/listing/$ITEM_ID" \
  -H "Content-Type: application/json" \
  -H "Authorization: $TOKEN" \
  -d '{"title":"Updated Test Item","description":"An updated test item"}' \
  | jq '.')
echo "$UPDATE_RESPONSE"

# Delete the listing with JWT auth
echo -e "\n7. Deleting the listing:"
curl -s -X DELETE "$BASE_URL/listing/$ITEM_ID" \
  -H "Authorization: $TOKEN" \
  -I
