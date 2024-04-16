package gql

import (
	"go-shop/graph"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

type GqlGenServer struct {
	port string
}

type GqlServer interface {
	Run()
}

func New(port string) *GqlGenServer {
	return &GqlGenServer{port}
}

func (s *GqlGenServer) Run() error {
	srv := createServer()

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", s.port)
	
	return http.ListenAndServe(":"+s.port, nil)
}

func createServer() *handler.Server {
	return handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{Resolvers: &graph.Resolver{}},
		),
	)
}
