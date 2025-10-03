# SmartCampus Workbench — Structure and Usage

This document explains the repository layout, how to run and build the backend and frontend, how to apply migrations, and where CI and metrics are located. It includes both Windows PowerShell and Linux commands.

## Repository layout (important files)

- `backend/` — Go backend service
  - `cmd/api/main.go` — server entrypoint
  - `internal/` — application internals
    - `handlers/` — HTTP handlers (auth, schools, assignments)
    - `db/` — DB connection helper (`Connect` with pooling & retries)
    - `middleware/metrics.go` — Prometheus instrumentation middleware
    - `auth/` — Casbin enforcer and RBAC middleware
    - `models/` — GORM models (User, School, Course, Assignment, ...)
    - `utils/validator.go` — request validation wrapper
  - `pkg/response/response.go` — unified JSON response helpers
  - `go.mod` / `go.sum` — Go modules
  - `Dockerfile` — production image build (if present)
  - `Makefile` and `scripts/` — helper build scripts

- `frontend/` — React + TypeScript frontend (Vite)
  - `src/` — source code (pages, services, ApiClient)
  - `src/services/apiClient.ts` — typed Axios wrapper used by pages
  - `package.json` / `tsconfig.json` — Node project config
  - `Dockerfile` — static site image builder

- `migrations/` — SQL migration files (e.g., `001_init.sql`)
- `scripts/` — helpful scripts
  - `migrate.sh` — applies SQL files to `DATABASE_DSN` using `psql`
  - `run-dev.sh` / `run-dev.ps1` — run backend and frontend for development
  - `scripts/linux/*` — Linux build scripts

- `docker-compose.yml` — compose to run backend, frontend, postgres (dev)
- `.github/workflows/ci.yml` — CI workflow to build backend and frontend; optional docker push
- `docs/` — generated documentation (this file)

## Quick start — prerequisites

- Go 1.21+ (for backend)
- Node 18+ / npm (for frontend)
- Docker & docker-compose (optional)
- PostgreSQL `psql` client (for migrations)

## Running locally (development)

### Windows (PowerShell)

Start backend and frontend dev servers concurrently (PowerShell):

```powershell
# from repo root
cd .\backend
# Backend dev: requires Go installed
go run ./cmd/api
# In a separate PowerShell
cd .\frontend
npm install
npm run dev
```

Or use the provided helper (runs both):

```powershell
# runs background processes using PowerShell
.\scripts\run-dev.ps1
```

### Linux / macOS (bash)

```bash
# run backend
cd backend
go run ./cmd/api &
# run frontend
cd frontend
npm ci
npm run dev &
wait
```

Or use the helper script:

```bash
./scripts/run-dev.sh
```

## Building production artifacts

### Backend (Linux binary)

PowerShell:

```powershell
cd .\backend
$env:CGO_ENABLED = "0"; $env:GOOS = "linux"; go build -a -installsuffix cgo -o main ./cmd/api
```

Linux:

```bash
cd backend
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api
```

### Frontend (Vite build)

PowerShell:

```powershell
cd .\frontend
npm ci
npm run build
```

Linux:

```bash
cd frontend
npm ci
npm run build
```

## Database migrations

Migrations are plain SQL files in `migrations/`.

Apply them with `psql` using environment variable `DATABASE_DSN`.

PowerShell example:

```powershell
$env:DATABASE_DSN = 'postgres://postgres:postgres@localhost:5432/smartcampus?sslmode=disable'
.\scripts\migrate.sh
```

Linux example:

```bash
export DATABASE_DSN='postgres://postgres:postgres@localhost:5432/smartcampus?sslmode=disable'
./scripts/migrate.sh
```

Notes:
- `scripts/migrate.sh` iterates over `migrations/*.sql` and applies them with `psql`.
- Ensure the database accepts connections from the host where you run migrations.

## Metrics and monitoring

- Prometheus metrics are exposed at `/metrics` on the backend server.
- Instrumentation middleware is in `backend/internal/middleware/metrics.go`:
  - Records `http_requests_total{method, path, status}` and `http_request_duration_seconds{method, path}`.
  - Path labels are sanitized to replace ID-like segments with `:id` to avoid high-cardinality labels.
  - Latency histogram uses tuned buckets [0.005, 0.01, 0.025, 0.05, 0.1, 0.3, 1.2, 5.0].

Operational tips:
- Ensure Prometheus scrapes the backend `/metrics` endpoint. Configure relabeling if you need different label names or to limit endpoints.
- For high-traffic endpoints, consider removing `path` label or aggregating to avoid cardinality explosion.

## CI

- GitHub Actions workflow is in `.github/workflows/ci.yml` and builds the backend and frontend on push/PR to `main`.
- The workflow optionally pushes Docker images if `DOCKER_USERNAME` and `DOCKER_PASSWORD` secrets are configured.
- The workflow uploads build artifacts (`backend/main`, `frontend/dist`) as artifacts for later use.

## Notable scripts

- `scripts/migrate.sh` — applies migrations using `psql` and `DATABASE_DSN`.
- `scripts/run-dev.sh` / `scripts/run-dev.ps1` — start both dev servers.
- `scripts/linux/build-backend.sh` and `scripts/linux/build-frontend.sh` — automated build scripts for CI/linux.

## Security & secrets

- Set `JWT_SECRET` (or `jwt.secret` in backend config) to a secure value in production.
- Use environment variables or a secrets manager for database credentials.

## Troubleshooting

- TypeScript build errors in frontend: run `npm ci` in `frontend/` and ensure the added dev types are installed (`@types/node`, etc.).
- If backend cannot connect to DB on startup, the DB connection uses retries; ensure `DATABASE_DSN` is correct and the DB is reachable.

## Next steps & suggestions

- Add a migration tool (golang-migrate) for versioned migrations and rollbacks.
- Add small integration tests to run migrations and smoke-test the API during CI.
- Consider adding a metrics dashboard (Grafana) and instrumenting business metrics (e.g., assignments created).

---

If you want, I can also update the root `README.md` with a short quickstart that links to this document and run a local frontend build to verify the TypeScript fixes. Which should I do next?
