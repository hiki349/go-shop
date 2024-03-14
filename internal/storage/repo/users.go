package repo

import (
	"context"
	"go-shop/internal/domain/models"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgxutil"
)

func (r Repo) FindUsers(ctx context.Context) ([]models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	query := "Select id, username, email, password, created_at, updated_at FROM users;"

	products, err := pgxutil.Select(ctx, r.db.Postgres, query, nil, pgx.RowToStructByPos[models.User])
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r Repo) FindUserByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	query := `Select
	id, username, email, password, created_at, updated_at
	FROM users
	WHERE id = $1;`

	product, err := pgxutil.SelectRow(ctx, r.db.Postgres, query, []any{id}, pgx.RowToStructByPos[models.User])
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *Repo) CreateUser(ctx context.Context, values models.User) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	sql := `INSERT INTO users
	(id, username, email, password, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6);`

	userID := uuid.New()
	_, err := pgxutil.ExecRow(
		ctx, r.db.Postgres, sql,
		userID,
		values.Username,
		values.Email,
		values.Password,
		time.Now(),
		nil,
	)
	if err != nil {
		return uuid.UUID{}, err
	}

	return userID, nil
}

func (r *Repo) UpdateUser(ctx context.Context, values models.User) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	sql := `UPDATE users
	SET username = $2, email = $3, password = $4, updated_at = $5
	WHERE id = $1;`

	_, err := pgxutil.ExecRow(
		ctx, r.db.Postgres, sql,
		values.ID,
		values.Username,
		values.Email,
		values.Password,
		time.Now(),
	)
	if err != nil {
		return values.ID, err
	}

	return values.ID, nil
}

func (r *Repo) DeleteUser(ctx context.Context, id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	sql := "DELETE FROM users WHERE id = $1;"

	_, err := pgxutil.ExecRow(ctx, r.db.Postgres, sql, id)
	if err != nil {
		return err
	}

	return nil
}
