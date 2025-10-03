#!/usr/bin/env bash
set -euo pipefail

echo "Building backend (linux)"
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o backend/main ./backend/cmd/api
echo "Backend built: backend/main"
