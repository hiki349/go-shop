package rest

import (
	"go-shop/internal/api/rest/handlers"
	"go-shop/internal/domain/services"
	"net/http"
)

type Rest struct {
	port string
	svc  *services.UsersService
}

func Init(port string, svc *services.UsersService) *Rest {
	return &Rest{
		port: port,
		svc:  svc,
	}
}

func (r *Rest) ServeHTTP() error {
	handlers.Init(r.svc).CreateRouter()

	return http.ListenAndServe(r.port, nil)
}
