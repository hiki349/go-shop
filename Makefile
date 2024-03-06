run: build
	@./bin/go-shop

build:
	@go build -o bin/go-shop cmd/main/main.go