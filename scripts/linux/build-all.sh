#!/usr/bin/env bash
set -euo pipefail

./scripts/linux/build-backend.sh
./scripts/linux/build-frontend.sh

echo "Build complete"
