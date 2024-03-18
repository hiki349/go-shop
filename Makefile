run: build
	@./bin/go-shop

db-up:
	@docker-compose up -d 

db-down:
	@docker-compose down

mgr-up:
	@goose -dir migrations postgres "postgresql://postgres:postgres@127.0.0.1:5432/shop-db" up

mgr-down:
	@goose -dir migrations postgres "postgresql://postgres:postgres@127.0.0.1:5432/shop-db" down

build:
	@go build -o bin/go-shop cmd/main/main.go

generate:
	@go run github.com/99designs/gqlgen generate.
