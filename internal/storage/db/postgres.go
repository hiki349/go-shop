package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Postgres struct {
	*pgx.Conn
}

func NewPostgres(ctx context.Context, connStr string) (*Postgres, error) {
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, err
	}

	return &Postgres{Conn: conn}, nil
}
