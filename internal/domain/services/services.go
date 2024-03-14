package services

import "go-shop/internal/storage/repo"

type Services struct {
	repo *repo.Repo
}

func New(repo *repo.Repo) *Services {
	return &Services{
		repo: repo,
	}
}
