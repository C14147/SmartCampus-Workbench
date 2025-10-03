#!/usr/bin/env bash
set -euo pipefail

if [ -z "${DATABASE_DSN:-}" ]; then
  echo "DATABASE_DSN is not set. Example: export DATABASE_DSN='postgres://postgres:postgres@localhost:5432/smartcampus?sslmode=disable'"
  exit 1
fi

echo "Applying migrations from ./migrations"

if command -v psql >/dev/null 2>&1; then
  for f in migrations/*.sql; do
    echo "Applying $f"
    psql "$DATABASE_DSN" -f "$f"
  done
  echo "Migrations applied"
else
  echo "psql not found. Please install psql or run the SQL files manually against your database."
  exit 1
fi
