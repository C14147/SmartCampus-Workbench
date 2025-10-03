# Developer Guide

This guide explains the minimal scaffold included in this workspace.

Backend (Go):
- Entry: `backend/cmd/api/main.go`
- Config: `backend/config/config.yaml.example` (viper)
- Handlers: `backend/internal/handlers/auth.go`
- Run: `go run ./cmd/api` (requires Go 1.21)

Frontend (Vite + React):
- Entry: `frontend/src/main.tsx`
- Dev: `cd frontend && npm install && npm run dev`

Docker compose:
- `docker-compose up --build` will build both images and expose ports 8080 (backend) and 3000 (frontend)

Next steps:
- Add persistent storage, migrations, and real auth
- Expand frontend with Fluent UI and Redux slices
