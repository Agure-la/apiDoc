# API Documentation Backend (Go)

## Overview

This repository contains the backend service for an API documentation application, implemented in **Go**.

The backend is responsible for **loading, validating, versioning, and serving API specifications** (primarily OpenAPI) in a structured, frontend-consumable format. It acts as the system of record for API documentation data and enforces consistency across versions.

The service is intentionally stateless, spec-driven, and optimized for developer experience.

---

## Responsibilities

The backend system handles:

- Loading OpenAPI specifications from configured sources
- Validating specs against the OpenAPI standard
- Normalizing and resolving schemas, paths, and references
- Supporting multiple API versions concurrently
- Exposing structured documentation data via HTTP APIs
- Optionally enabling interactive request execution (“try-it-out”)
- Providing operational endpoints (health, readiness)

The backend does **not** render UI and does **not** act as an API gateway.

---

## Architecture

High-level flow:

OpenAPI Specs ──▶ Loader ──▶ Parser ──▶ Normalizer ──▶ In-Memory Cache
│
▼
HTTP JSON API
│
▼
Docs Frontend


The backend translates raw OpenAPI documents into stable, frontend-friendly domain models.

---

## Project Structure

api-docs/
├── cmd/
│ └── server/
│ └── main.go
├── internal/
│ ├── http/
│ │ ├── server.go
│ │ ├── routes.go
│ │ └── handlers/
│ │ ├── health.go
│ │ └── api_docs.go
│ ├── spec/
│ │ ├── loader.go
│ │ ├── parser.go
│ │ ├── normalizer.go
│ │ └── model.go
│ ├── versioning/
│ │ └── resolver.go
│ └── config/
│ └── config.go
├── go.mod
└── README.md


### Structure Rationale

- `cmd/` contains application entry points only
- `internal/` prevents accidental external reuse
- HTTP layer is thin and delegates to domain logic
- Spec processing is isolated from transport concerns

---

## HTTP API

All endpoints return JSON and are intended to be consumed by a documentation frontend.

### Health

GET /health


Response:
```json
{
  "status": "ok"
}
APIs
GET /apis
Returns a list of available APIs.

GET /apis/{api}
Returns metadata for a specific API.

Versions
GET /apis/{api}/versions
Returns all known versions for an API.

GET /apis/{api}/versions/{version}
Returns documentation metadata for a specific API version.