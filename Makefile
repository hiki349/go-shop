dev:
	@go run cmd/main/main.go

run: build
	@./bin/go-shop

app-up:
	@docker compose up -d

dev-up:
	@docker compose -f ./docker-compose-dev.yml up -d

app-down:
	@docker compose down

mgr-up:
	@goose -dir migrations postgres "postgresql://postgres:postgres@127.0.0.1:5432/shop-db" up

mgr-down:
	@goose -dir migrations postgres "postgresql://postgres:postgres@127.0.0.1:5432/shop-db" down

build:
	@go build -o bin/go-shop -ldflags "-s -w" -a cmd/main/main.go

generate:
	@go run github.com/99designs/gqlgen generate.
