param()

Write-Host "Building backend..."
env:CGO_ENABLED = "0"
env:GOOS = "linux"
go build -a -installsuffix cgo -o main ./cmd/api
Write-Host "Built: main"
