package gql

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"go-shop/internal/api/gql/generated"
	"go-shop/internal/api/gql/resolvers"
	"go-shop/internal/domain/services"
	"go-shop/internal/middleware"
)

type GqlGenServer struct {
	port            string
	ProductsService *services.ProductsService
	UsersService    *services.UsersService
	CartsService    *services.CartsService
}

type GqlServer interface {
	Run()
}

func New(port string, productsService *services.ProductsService, usersService *services.UsersService, cartsService *services.CartsService) *GqlGenServer {
	return &GqlGenServer{
		port:            port,
		ProductsService: productsService,
		UsersService:    usersService,
		CartsService:    cartsService,
	}
}

func (s *GqlGenServer) Run() error {
	logging := middleware.Auth()
	srv := s.createServer()

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", logging(srv))
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", s.port)

	return http.ListenAndServe(":"+s.port, nil)
}

func (s *GqlGenServer) createServer() *handler.Server {
	return handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &resolvers.Resolver{
				ProductsService: s.ProductsService,
				UsersService:    s.UsersService,
				CartsService:    s.CartsService,
			}},
		),
	)
}
