package db

import (
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinIO struct {
	client *minio.Client
}

func NewMinIO(ctx context.Context, ssl bool, url, user, password string) (*MinIO, error) {
	client, err := minio.New(url, &minio.Options{
		Creds:  credentials.NewStaticV4(user, password, ""),
		Secure: ssl,
	})
	if err != nil {
		return nil, err
	}

	return &MinIO{client: client}, nil
}
