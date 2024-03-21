package handlers

import (
	"go-shop/internal/domain/services"
	"net/http"
)

type Handler struct {
	svc *services.UsersService
}

func Init(svc *services.UsersService) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) CreateRouter() {
	http.HandleFunc("POST /login", h.login)
	http.HandleFunc("POST /logout", h.logout)
}
