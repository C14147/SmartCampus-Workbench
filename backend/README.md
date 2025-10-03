# SmartCampus Backend (minimal scaffold)

This folder contains a minimal Go backend scaffold for the SmartCampus project.

Features included:
- Gin HTTP server with health endpoint
- Minimal JWT-based login stub (returns a token for demo user)
- Config loader via Viper
- Dockerfile for building the server image

How to run locally (requires Go 1.21):

1. Copy the example config if you want to set a secret:

   cp config/config.yaml.example config/config.yaml

2. (Optional) Configure PostgreSQL and provide a DSN via env:

   $env:DATABASE_DSN = 'postgres://user:pass@localhost:5432/smartcampus?sslmode=disable'

3. Build and run:

   go run ./cmd/api

Endpoints:
- GET /health
- POST /api/v1/auth/login  { username, password }
- GET /api/v1/auth/me

This scaffold is intentionally small. Extend handlers, add persistent storage,
authentication middleware, and tests as next steps.
