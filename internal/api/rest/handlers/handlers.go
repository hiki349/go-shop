package handlers

import (
	"go-shop/internal/domain/services"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	svc    *services.AuthService
	router *chi.Mux
}

func Init(svc *services.AuthService, router *chi.Mux) *Handler {
	return &Handler{
		svc:    svc,
		router: router,
	}
}

func (h *Handler) CreateRouter() {
	h.router.Post("/login", h.login)
	h.router.Post("/logout", h.logout)
	h.router.Get("/token", h.getAccessToken)
}
