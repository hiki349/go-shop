package services

import (
	"context"
	"fmt"
	"go-shop/internal/pkg/token"
	"go-shop/internal/storage/repo"
)

type AuthService struct {
	repo   repo.UsersRepo
	secret string
}

func NewAuthService(repo repo.UsersRepo, secret string) *AuthService {
	return &AuthService{
		repo:   repo,
		secret: secret,
	}
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if user.Password != password {
		return "", fmt.Errorf("invalid password or email")
	}
	refreshTokenString, err := token.CreateRefreshToken(user.ID, s.secret)

	return refreshTokenString, err
}

func (s *AuthService) GetAccessToken(ctx context.Context, refreshTokenString string) (string, error) {
	refreshToken, err := token.VerifyToken(refreshTokenString, s.secret)
	if err != nil {
		return "", fmt.Errorf("invalid refresh token")
	}
	accessTokenString, err := token.CreateAccessToken(refreshToken, s.secret)

	return accessTokenString, err
}
