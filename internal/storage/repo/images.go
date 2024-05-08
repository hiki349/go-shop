package repo

import (
	"context"

	"go-shop/internal/storage/db"
	"go-shop/internal/storage/models"
)

type MinIOImagesRepo struct {
	db *db.MinIO
}

type ImagesRepo interface {
	Save(ctx context.Context, image *models.Image) (string, error)
	Get(ctx context.Context, id string) (*models.Image, error)
}

func NewMinIOImagesRepo(db *db.MinIO) *MinIOImagesRepo {
	return &MinIOImagesRepo{db}
}

func (r *MinIOImagesRepo) Save(ctx context.Context, image *models.Image) (string, error) {
	return "", nil
}

func (r *MinIOImagesRepo) Get(ctx context.Context, id string) (*models.Image, error) {
	return nil, nil
}
