package services

import (
	"context"
	"go-shop/internal/domain/models"
	"go-shop/internal/pkg/token"
	"go-shop/internal/storage/repo"
)

type AuthService struct {
	repo repo.IUsersRepo
}

type AuthResult struct {
	Token        string
	RefreshToken string
	User         *models.User
	Err          error
}

func (svc AuthService) Login(ctx context.Context, email, password string) *AuthResult {
	user, err := svc.repo.FindUserByEmail(ctx, email)
	if err != nil {
		return &AuthResult{Err: err}
	}
	accessToken, err := token.NewToken("", 0, &token.TokenUserInfo{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	})
	if err != nil {
		return &AuthResult{Err: err}
	}

	refreshToken, err := token.NewToken("", 0, nil)
	if err != nil {
		return &AuthResult{Err: err}
	}

	return &AuthResult{
		Token:        accessToken,
		RefreshToken: refreshToken,
		User:         user,
		Err:          nil,
	}
}
