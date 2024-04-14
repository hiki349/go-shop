package handlers

import (
	"context"
	"encoding/json"
	"go-shop/internal/api/rest/model"
	"go-shop/internal/pkg/cookie"
	"net/http"
)

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	var user model.UserLogin
	ctx := context.Background()

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.svc.Login(ctx, user.Email, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	cookie.SetCookie("refresh_token", token, w)
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) logout(w http.ResponseWriter, r *http.Request) {
	cookie.ClearCookie("refresh_token", w)

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) getAccessToken(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	refreshToken, err := cookie.GetCookieValue("refresh_token", r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	accesToken, err := h.svc.GetAccessToken(ctx, refreshToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	jsonBytes, err := json.Marshal(map[string]string{"access_token": accesToken})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(jsonBytes)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
