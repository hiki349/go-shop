package main

import (
	"context"
	"fmt"
	"go-shop/internal/api/gql/resolvers"
	"go-shop/internal/api/gql/runtime"
	"go-shop/internal/domain/services"
	"go-shop/internal/storage/db"
	"go-shop/internal/storage/repo"
	"log"

	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
)

func main() {
	config := mustGetConfig()

	db, err := db.New(context.Background(), config.connStr)
	if err != nil {
		log.Fatal(err)
	}

	repo := repo.New(db)
	svc := services.New(repo)

	log.Fatal(startGqlServer(svc, config.port))
}

func startGqlServer(svc *services.Services, port string) error {
	srv := handler.NewDefaultServer(
		runtime.NewExecutableSchema(
			runtime.Config{
				Resolvers: &resolvers.Resolver{Services: svc},
			},
		),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	return http.ListenAndServe(":"+port, nil)
}

type Config struct {
	port    string
	connStr string
}

func mustGetConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")

	dbUsername := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s", dbUsername, dbPass, dbPort, dbName)

	return Config{
		port:    port,
		connStr: connStr,
	}
}
