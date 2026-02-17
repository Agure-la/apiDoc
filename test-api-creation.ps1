# Base URL
$BaseUrl = "http://localhost:8080"

Write-Host "=== Refactored API Documentation Backend Test ===" -ForegroundColor Green
Write-Host ""

# Test health endpoint
Write-Host "1. Testing health endpoint..." -ForegroundColor Yellow
try {
    $response = Invoke-RestMethod -Uri "$BaseUrl/health" -Method Get
    $response | ConvertTo-Json -Depth 10
} catch {
    Write-Host "Error: $($_.Exception.Message)" -ForegroundColor Red
}
Write-Host ""

# Create a new API
Write-Host "2. Creating a new API..." -ForegroundColor Yellow
$createBody = @{
    name = "user-service"
    title = "User Service API"
    description = "API for managing users"
    version = "v1"
    spec = @{
        openapi = "3.0.0"
        info = @{
            title = "User Service API"
            version = "1.0.0"
        }
        paths = @{
            "/users" = @{
                get = @{
                    summary = "List all users"
                    responses = @{
                        "200" = @{
                            description = "List of users"
                        }
                    }
                }
            }
        }
    }
    metadata = @{
        team = "backend"
        repository = "github.com/company/user-service"
    }
} | ConvertTo-Json -Depth 10

try {
    $response = Invoke-RestMethod -Uri "$BaseUrl/apis" -Method Post -ContentType "application/json" -Body $createBody
    $response | ConvertTo-Json -Depth 10
} catch {
    Write-Host "Error: $($_.Exception.Message)" -ForegroundColor Red
}
Write-Host ""

# List all APIs
Write-Host "3. Listing all APIs..." -ForegroundColor Yellow
try {
    $response = Invoke-RestMethod -Uri "$BaseUrl/apis" -Method Get
    $response | ConvertTo-Json -Depth 10
} catch {
    Write-Host "Error: $($_.Exception.Message)" -ForegroundColor Red
}
Write-Host ""

# Get specific API
Write-Host "4. Getting specific API..." -ForegroundColor Yellow
try {
    $response = Invoke-RestMethod -Uri "$BaseUrl/apis/user-service" -Method Get
    $response | ConvertTo-Json -Depth 10
} catch {
    Write-Host "Error: $($_.Exception.Message)" -ForegroundColor Red
}
Write-Host ""

# Update the API
Write-Host "5. Updating the API..." -ForegroundColor Yellow
$updateBody = @{
    title = "User Service API (Updated)"
    description = "API for managing users - updated version"
    metadata = @{
        team = "backend"
        repository = "github.com/company/user-service"
        status = "active"
    }
} | ConvertTo-Json -Depth 10

try {
    $response = Invoke-RestMethod -Uri "$BaseUrl/apis/user-service" -Method Put -ContentType "application/json" -Body $updateBody
    $response | ConvertTo-Json -Depth 10
} catch {
    Write-Host "Error: $($_.Exception.Message)" -ForegroundColor Red
}
Write-Host ""

# Get updated API
Write-Host "6. Getting updated API..." -ForegroundColor Yellow
try {
    $response = Invoke-RestMethod -Uri "$BaseUrl/apis/user-service" -Method Get
    $response | ConvertTo-Json -Depth 10
} catch {
    Write-Host "Error: $($_.Exception.Message)" -ForegroundColor Red
}
Write-Host ""

# Delete the API
Write-Host "7. Deleting the API..." -ForegroundColor Yellow
try {
    $response = Invoke-RestMethod -Uri "$BaseUrl/apis/user-service" -Method Delete
    $response | ConvertTo-Json -Depth 10
} catch {
    Write-Host "Error: $($_.Exception.Message)" -ForegroundColor Red
}
Write-Host ""

# List APIs after deletion
Write-Host "8. Listing APIs after deletion..." -ForegroundColor Yellow
try {
    $response = Invoke-RestMethod -Uri "$BaseUrl/apis" -Method Get
    $response | ConvertTo-Json -Depth 10
} catch {
    Write-Host "Error: $($_.Exception.Message)" -ForegroundColor Red
}
Write-Host ""

Write-Host "=== Refactoring Test Complete ===" -ForegroundColor Green
