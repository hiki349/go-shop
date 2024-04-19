package gql

import (
	"go-shop/graph"
	"go-shop/internal/domain/services"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

type GqlGenServer struct {
	port            string
	ProductsService *services.ProductsService
}

type GqlServer interface {
	Run()
}

func New(port string, productsService *services.ProductsService) *GqlGenServer {
	return &GqlGenServer{
		port:            port,
		ProductsService: productsService,
	}
}

func (s *GqlGenServer) Run() error {
	srv := createServer(s.ProductsService)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", s.port)

	return http.ListenAndServe(":"+s.port, nil)
}

func createServer(productsService *services.ProductsService) *handler.Server {
	return handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{Resolvers: &graph.Resolver{ProductsService: productsService}},
		),
	)
}
