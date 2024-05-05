package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	*pgxpool.Pool
}

var maxAttempts = 3

func NewPostgres(ctx context.Context, connStr string) (pg *Postgres, err error) {
	for maxAttempts > 0 {
		ctx, cancel := context.WithTimeout(ctx, time.Second*5)
		defer cancel()

		pg.Pool, err = pgxpool.New(ctx, connStr)
		if err != nil {
			maxAttempts--
			continue
		}

		break
	}

	return pg, err
}
