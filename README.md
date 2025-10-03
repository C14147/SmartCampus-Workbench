# SmartCampus Workbench

This repository is a minimal, opinionated scaffold for the SmartCampus project. It contains a lightweight backend (Go + Gin) and frontend (React + Vite), Dockerfiles, and helper scripts to get started.

See detailed structure and usage in `docs/STRUCTURE_AND_USAGE.md`.

Quick start (Windows PowerShell):

1. Copy env example:

   copy .env.example .env

2. Start with docker-compose (Docker required):

   docker-compose up --build

Or run dev servers locally (requires Go and Node):

   # PowerShell
   .\scripts\run-dev.ps1

   # bash
   ./scripts/run-dev.sh

See `docs/STRUCTURE_AND_USAGE.md` for full instructions and operational notes.
