env:
	@cp .env.example .env

install:
	@go mod download

build:
	@go build -o bin/feedks-api ./cmd/main.go

run: build
	@./bin/feedks-api
