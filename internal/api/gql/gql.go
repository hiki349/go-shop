package gql

import (
	"go-shop/internal/api/gql/resolvers"
	"go-shop/internal/api/gql/runtime"
	"go-shop/internal/domain/services"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func MustStartGqlServer(
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
