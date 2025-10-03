param()

# Simple script to build backend and run frontend dev server.
Write-Host "Starting backend (go run) and frontend (npm run dev)"

Start-Process -NoNewWindow -FilePath pwsh -ArgumentList '-Command', 'cd backend; go run ./cmd/api' -WindowStyle Hidden
Start-Process -NoNewWindow -FilePath pwsh -ArgumentList '-Command', 'cd frontend; npm run dev' -WindowStyle Hidden
