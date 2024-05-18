package repo

import (
	"context"
	"errors"

	"go-shop/internal/storage/db"
)

var (
	ErrTokenNotFound = errors.New("not found token")
)

type MongoTokensRepo struct {
	client *db.Mongo
}

type TokensRepo interface {
	Exists(token string) error
	Add(token string) error
}

func NewMongoTokensRepo(client *db.Mongo) *MongoTokensRepo {
	return &MongoTokensRepo{
		client: client,
	}
}

func (r *MongoTokensRepo) Exists(ctx context.Context, token string) error {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	collection := r.client.Database("tokens").Collection("tokens")

	res := collection.FindOne(ctx, token)
	if err := res.Err(); err != nil {
		return err
	}

	return nil
}

func (r *MongoTokensRepo) Add(ctx context.Context, token string) error {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	collection := r.client.Database("tokens").Collection("tokens")

	_, err := collection.InsertOne(ctx, token)
	if err != nil {
		return err
	}

	return nil
}
