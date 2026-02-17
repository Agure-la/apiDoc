#!/bin/bash

# Base URL
BASE_URL="http://localhost:8080"

echo "=== API Documentation Backend Test ==="
echo

# Test health endpoint
echo "1. Testing health endpoint..."
curl -s "$BASE_URL/health" | jq .
echo

# Create a new API
echo "2. Creating a new API..."
CREATE_RESPONSE=$(curl -s -X POST "$BASE_URL/apis" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "user-service",
    "title": "User Service API",
    "description": "API for managing users",
    "version": "v1",
    "spec": {
      "openapi": "3.0.0",
      "info": {
        "title": "User Service API",
        "version": "1.0.0"
      },
      "paths": {
        "/users": {
          "get": {
            "summary": "List all users",
            "responses": {
              "200": {
                "description": "List of users"
              }
            }
          }
        }
      }
    },
    "metadata": {
      "team": "backend",
      "repository": "github.com/company/user-service"
    }
  }')

echo "$CREATE_RESPONSE" | jq .
echo

# List all APIs
echo "3. Listing all APIs..."
curl -s "$BASE_URL/apis" | jq .
echo

# Get specific API
echo "4. Getting specific API..."
curl -s "$BASE_URL/apis/user-service" | jq .
echo

# Update the API
echo "5. Updating the API..."
curl -s -X PUT "$BASE_URL/apis/user-service" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "User Service API (Updated)",
    "description": "API for managing users - updated version",
    "metadata": {
      "team": "backend",
      "repository": "github.com/company/user-service",
      "status": "active"
    }
  }' | jq .
echo

# Get updated API
echo "6. Getting updated API..."
curl -s "$BASE_URL/apis/user-service" | jq .
echo

# Delete the API
echo "7. Deleting the API..."
curl -s -X DELETE "$BASE_URL/apis/user-service" | jq .
echo

# List APIs after deletion
echo "8. Listing APIs after deletion..."
curl -s "$BASE_URL/apis" | jq .
echo

echo "=== Test Complete ==="
