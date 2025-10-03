#!/usr/bin/env bash
set -e

echo "Starting backend and frontend dev servers"
(cd backend && go run ./cmd/api) &
(cd frontend && npm run dev) &
wait
