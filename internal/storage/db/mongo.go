package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	*mongo.Client
}

func NewMongo(ctx context.Context, connStr string) (*Mongo, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connStr))
	if err != nil {
		return nil, err
	}

	return &Mongo{Client: client}, nil
}
