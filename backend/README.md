# SmartCampus Workbench Backend

A Go-based backend for SmartCampus, using Gin, GORM, PostgreSQL, and Redis.

## Features
- JWT authentication & RBAC
- User, course, assignment, submission, message, and file management
- PostgreSQL for persistent storage
- Redis for caching and session management
- Modular architecture (handlers, services, repositories)
- Docker & docker-compose deployment

## Project Structure
```
backend/
├── cmd/api/main.go                # Entry point
├── internal/
│   ├── config/config.go           # Config management
│   ├── database/postgres.go       # PostgreSQL connection
│   ├── database/redis.go          # Redis connection
│   ├── models/                    # Data models
│   ├── handlers/                  # HTTP handlers
│   ├── services/                  # Business logic
│   ├── repositories/              # Data access
│   ├── middleware/                # Gin middlewares
│   ├── utils/                     # Utility functions
│   └── websocket/                 # WebSocket support
├── pkg/
│   ├── response/response.go       # Unified response
│   └── cache/                     # Cache interface & Redis impl
├── scripts/init_db.sql            # DB schema
├── Dockerfile                     # Backend Dockerfile
├── deploy.sh, stop.sh             # Deployment scripts
├── API_DOCUMENT.md                # API documentation
└── README.md                      # This file
```

## Quick Start

### 1. Build & Run (Docker)
```bash
cd backend
./deploy.sh
```

### 2. Stop
```bash
./stop.sh
```

### 3. Manual DB Init
If not using docker-compose, run:
```bash
psql -U postgres -d smartcampus -f scripts/init_db.sql
```

### 4. Environment Variables
Set in docker-compose or `.env`:
- `DB_HOST`, `DB_PORT`, `DB_NAME`, `DB_USER`, `DB_PASSWORD`
- `REDIS_URL`
- `JWT_SECRET`

## API Documentation
See [API_DOCUMENT.md](./API_DOCUMENT.md)

## Development
- Go 1.21+
- PostgreSQL 15+
- Redis 7+
- Install dependencies: `go mod tidy`
- Run locally: `go run cmd/api/main.go`

