# Windows Setup Script for SmartCampus Workbench (Based on plan.md)
# Save as setup_plan.ps1 and run in PowerShell (Run as Administrator if needed)

Write-Host "=== SmartCampus Workbench Setup (plan.md) ===" -ForegroundColor Cyan

# 1. Prerequisites
Write-Host "Checking Docker installation..." -ForegroundColor Yellow
docker --version
docker-compose --version

# 2. Clone Repository
Write-Host "Cloning repository..." -ForegroundColor Yellow
git clone <your-repo-url>
cd SmartCampus-Workbench

# 3. Environment Configuration
Write-Host "Copying environment file..." -ForegroundColor Yellow
Copy-Item .env.example .env
Write-Host "Please edit .env with your settings." -ForegroundColor Yellow
notepad .env

# 4. Build and Start Services
Write-Host "Building Docker services..." -ForegroundColor Yellow
docker-compose build
Write-Host "Starting Docker services..." -ForegroundColor Yellow
docker-compose up -d
docker-compose ps
docker-compose logs -f

# 5. Database Setup
Write-Host "Running database migrations..." -ForegroundColor Yellow
docker-compose exec backend npx prisma migrate deploy
Write-Host "Seeding initial data (optional)..." -ForegroundColor Yellow
docker-compose exec backend npm run db:seed

# 6. Health Checks
Write-Host "Checking health endpoints..." -ForegroundColor Yellow
Invoke-WebRequest -Uri "http://localhost:3001/health" -UseBasicParsing
Invoke-WebRequest -Uri "http://localhost/" -UseBasicParsing
docker-compose exec postgres pg_isready -U postgres
docker-compose exec redis redis-cli ping

Write-Host "=== Setup Complete ===" -ForegroundColor Green
Write-Host "Access Frontend: http://localhost" -ForegroundColor White
Write-Host "Access Backend API: http://localhost:3001" -ForegroundColor White
Write-Host "API Docs: http://localhost:3001/api/docs" -ForegroundColor White
