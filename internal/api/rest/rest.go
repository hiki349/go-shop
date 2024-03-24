package rest

import (
	"go-shop/internal/api/rest/handlers"
	"go-shop/internal/domain/services"
	"log"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Rest struct {
	port string
	svc  *services.UsersService
	clog *slog.Logger
}

func MustStartRestServer(svc *services.UsersService, port string, clog *slog.Logger) {
	srv := &Rest{
		port: port,
		svc:  svc,
		clog: clog,
	}

	log.Printf("connect to http://localhost%s/ for rest server", port)
	log.Fatal(srv.ServeHTTP())
}

func (r *Rest) ServeHTTP() error {
	router := chi.NewRouter()
	handlers.Init(r.svc, router).CreateRouter()

	return http.ListenAndServe(r.port, router)
}
