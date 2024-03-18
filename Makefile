run: build
	@./bin/go-shop

mgr-up:
	@docker-compose up -d

mgr-down:
	@docker-compose down

build:
	@go build -o bin/go-shop cmd/main/main.go

generate:
	@go run github.com/99designs/gqlgen generate.