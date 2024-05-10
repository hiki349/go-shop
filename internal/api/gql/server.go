package gql

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"go-shop/internal/api/gql/generated"
	"go-shop/internal/api/gql/resolvers"
	"go-shop/internal/domain/services"
)

type GqlGenServer struct {
	port            string
	ProductsService *services.ProductsService
	UsersService    *services.UsersService
}

type GqlServer interface {
	Run()
}

func New(port string, productsService *services.ProductsService, usersService *services.UsersService) *GqlGenServer {
	return &GqlGenServer{
		port:            port,
		ProductsService: productsService,
		UsersService:    usersService,
	}
}

func (s *GqlGenServer) Run() error {
	srv := createServer(s.ProductsService, s.UsersService)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", s.port)

	return http.ListenAndServe(":"+s.port, nil)
}

func createServer(productsService *services.ProductsService, usersService *services.UsersService) *handler.Server {
	return handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &resolvers.Resolver{
				ProductsService: productsService,
				UsersService:    usersService,
			}},
		),
	)
}
