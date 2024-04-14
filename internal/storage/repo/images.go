package repo

import "go-shop/internal/storage/db"

type MinIOImagesRepo struct {
	db *db.MinIO
}

type ImagesRepo interface{}

func NewMinIOImagesRepo(db *db.MinIO) *MinIOImagesRepo {
	return &MinIOImagesRepo{db}
}
