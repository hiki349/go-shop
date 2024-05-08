package services

import (
	"context"

	"go-shop/internal/storage/repo"
	"go-shop/pkg/token"
)

type TokensService struct {
	repo   repo.TokensRepo
	secret string
}

type ITokensService interface {
	GetAccessToken(ctx context.Context, refreshToken string) (string, error)
	Logout(ctx context.Context, refreshToken string) error
}

func (s *TokensService) GetAccessToken(ctx context.Context, refreshToken string) (string, error) {
	err := s.repo.Exists(refreshToken)
	if err != nil {
		return "", err
	}

	refresh, err := token.VerifyToken(refreshToken, s.secret)
	if err != nil {
		return "", err
	}

	return token.CreateAccessToken(refresh, s.secret)
}

func (s *TokensService) Logout(ctx context.Context, refreshToken string) error {
	err := s.repo.Exists(refreshToken)
	if err != nil {
		return err
	}

	_, err = token.VerifyToken(refreshToken, s.secret)
	if err != nil {
		return err
	}

	return s.repo.Add(refreshToken)
}
