#!/usr/bin/env bash
set -euo pipefail

echo "Building frontend"
cd frontend
npm ci
npm run build
cd -
echo "Frontend built: frontend/dist"
