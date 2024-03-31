package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	Postgres *pgx.Conn
	Mongo    *mongo.Client
}

func New(ctx context.Context, connStr string) (*DB, error) {
	postgres, err := newPostgres(ctx, connStr)
	if err != nil {
		return nil, err
	}

	mongo, err := newMongo(ctx, connStr)
	if err != nil {
		return nil, err
	}

	return &DB{
		Postgres: postgres,
		Mongo:    mongo,
	}, nil
}

func newPostgres(ctx context.Context, connStr string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func newMongo(ctx context.Context, connStr string) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connStr))
	if err != nil {
		return nil, err
	}

	return client, nil
}
