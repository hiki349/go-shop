package main

import (
	"context"
	"go-shop/configuration"
	"go-shop/internal/api/gql/resolvers"
	"go-shop/internal/api/gql/runtime"
	"go-shop/internal/api/rest"
	"go-shop/internal/domain/services"
	"go-shop/internal/pkg/logger"
	"go-shop/internal/storage/db"
	"go-shop/internal/storage/repo"
	"log"

	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	config := configuration.MustGetConfig()

	clog := logger.New(config.Mode)
	clog.Info("Hello world")

	db, err := db.New(context.Background(), config.ConnStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Postgres.Close(context.Background())

	productsRepo := repo.NewProductsRepo(db)
	cartsRepo := repo.NewCartsRepo(db)
	userssRepo := repo.NewUsersRepo(db)

	productsService := services.NewProductsService(productsRepo)
	cartsService := services.NewCartsService(cartsRepo)
	usersService := services.NewUsersService(userssRepo)

	go rest.MustStartRestServer(usersService, config.RestPort, clog)
	mustStartGqlServer(productsService, cartsService, usersService, config.GqlPort)
}

func mustStartGqlServer(
	productsService *services.ProductsService,
	cartsService *services.CartsService,
	usersService *services.UsersService,
	port string,
) {
	srv := handler.NewDefaultServer(
		runtime.NewExecutableSchema(
			runtime.Config{
				Resolvers: &resolvers.Resolver{
					ProductsService: productsService,
					CartsService:    cartsService,
					UsersService:    usersService,
				},
			},
		),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
