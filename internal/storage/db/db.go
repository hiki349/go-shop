package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type DB struct {
	Postgres *pgx.Conn
}

func New(ctx context.Context, connStr string) (*DB, error) {
	postgres, err := newPostgres(ctx, connStr)
	if err != nil {
		return nil, err
	}

	return &DB{
		Postgres: postgres,
	}, nil
}

func newPostgres(ctx context.Context, connStr string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, err
	}
	defer conn.Close(ctx)

	return conn, nil
}
