package services

import "go-shop/internal/storage/repo"

type ImagesService struct {
	repo repo.ImagesRepo
}

func NewImagesService(repo repo.ImagesRepo) *ImagesService {
	return &ImagesService{repo}
}
