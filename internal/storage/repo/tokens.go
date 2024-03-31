package repo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type TokensRepo struct {
	client *mongo.Client
}

type ITokensRepo interface {
	Exists(token string) (bool, error)
	Add(token string) error
}

func NewTokensRepo(client *mongo.Client) *TokensRepo {
	return &TokensRepo{
		client: client,
	}
}

func (r *TokensRepo) Exists(ctx context.Context, token string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	collection := r.client.Database("tokens").Collection("tokens")

	res := collection.FindOne(ctx, token)
	if err := res.Err(); err != nil {
		return false, err
	}

	return true, nil
}

func (r *TokensRepo) Add(ctx context.Context, token string) error {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	collection := r.client.Database("tokens").Collection("tokens")

	_, err := collection.InsertOne(ctx, token)
	if err != nil {
		return err
	}

	return nil
}
