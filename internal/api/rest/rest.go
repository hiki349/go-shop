package rest

import (
	"go-shop/internal/api/rest/handlers"
	"go-shop/internal/domain/services"
	"log"
	"log/slog"
	"net/http"
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

	log.Fatal(srv.ServeHTTP())
}

func (r *Rest) ServeHTTP() error {
	handlers.Init(r.svc).CreateRouter()

	return http.ListenAndServe(r.port, nil)
}
