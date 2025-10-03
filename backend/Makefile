BINARY=main

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $(BINARY) ./cmd/api

.PHONY: run
run:
	go run ./cmd/api
