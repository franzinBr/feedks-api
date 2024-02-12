build:
	@go build -o bin/feedks-api ./cmd/main.go

run: build
	@./bin/feedks-api
