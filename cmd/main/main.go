package main

import (
	"context"
	"go-shop/configuration"
	"go-shop/internal/api/gql/resolvers"
	"go-shop/internal/api/gql/runtime"
	"go-shop/internal/api/rest"
	"go-shop/internal/domain/services"
	"go-shop/internal/storage/db"
	"go-shop/internal/storage/repo"
	"log"

	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	config := configuration.MustGetConfig()

	db, err := db.New(context.Background(), config.ConnStr)
	if err != nil {
		log.Fatal(err)
	}

	repo := repo.New(db)
	svc := services.New(repo)

	log.Fatal(startGqlServer(svc, config.GqlPort))
	log.Fatal(startRestServer(svc, config.RestPort))
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

func startRestServer(svc *services.Services, port string) error {
	srv := rest.Init(port, svc)

	return srv.ServeHTTP()
}
