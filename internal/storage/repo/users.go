package repo

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgxutil"

	"go-shop/internal/domain/models"
	"go-shop/internal/storage/db"
)

var (
	ErrUserNotFound = errors.New("not found user")
)

type PostgresUsersRepo struct {
	db *db.Postgres
}

type UsersRepo interface {
	FindAll(ctx context.Context) ([]models.User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	Create(ctx context.Context, data *models.User) (uuid.UUID, error)
	Update(ctx context.Context, data *models.User) (uuid.UUID, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

func NewUsersRepo(db *db.Postgres) *PostgresUsersRepo {
	return &PostgresUsersRepo{db}
}

func (r PostgresUsersRepo) FindAll(ctx context.Context) ([]models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	query := "SELECT id, username, email, password, created_at, updated_at FROM users;"

	users, err := pgxutil.Select(ctx, r.db, query, nil, pgx.RowToStructByPos[models.User])
	if err != nil {
		slog.InfoContext(ctx, "%w", err)

		return nil, err
	}

	return users, nil
}

func (r PostgresUsersRepo) FindByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	query := `Select
	id, username, email, password, created_at, updated_at
	FROM users
	WHERE id = $1;`

	user, err := pgxutil.SelectRow(ctx, r.db, query, []any{id}, pgx.RowToStructByPos[models.User])
	if err != nil {
		slog.InfoContext(ctx, "%w", err)

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}

		return nil, err
	}

	return &user, nil
}

func (r PostgresUsersRepo) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	query := `SELECT
	id, username, email, password, created_at, updated_at
	FROM users
	WHERE email = $1;`

	user, err := pgxutil.SelectRow(ctx, r.db, query, []any{email}, pgx.RowToStructByPos[models.User])
	if err != nil {
		slog.InfoContext(ctx, "%w", err)

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}

		return nil, err
	}

	return &user, err
}

func (r *PostgresUsersRepo) Create(ctx context.Context, values *models.User) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	sql := `INSERT INTO users
	(id, username, email, password, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6);`

	userID := uuid.New()
	_, err := pgxutil.ExecRow(
		ctx, r.db, sql,
		userID,
		values.Username,
		values.Email,
		values.Password,
		time.Now(),
		nil,
	)
	if err != nil {
		slog.InfoContext(ctx, "%w", err)

		if errors.Is(err, pgx.ErrNoRows) {
			return uuid.UUID{}, ErrUserNotFound
		}

		return uuid.UUID{}, err
	}

	return userID, nil
}

func (r *PostgresUsersRepo) Update(ctx context.Context, values *models.User) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	sql := `UPDATE users
	SET username = $2, email = $3, password = $4, updated_at = $5
	WHERE id = $1;`

	_, err := pgxutil.ExecRow(
		ctx, r.db, sql,
		values.ID,
		values.Username,
		values.Email,
		values.Password,
		time.Now(),
	)
	if err != nil {
		slog.InfoContext(ctx, "%w", err)

		if errors.Is(err, pgx.ErrNoRows) {
			return uuid.UUID{}, ErrNotFound
		}

		return uuid.UUID{}, err
	}

	return values.ID, nil
}

func (r *PostgresUsersRepo) Delete(ctx context.Context, id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	sql := "DELETE FROM users WHERE id = $1;"

	_, err := pgxutil.ExecRow(ctx, r.db, sql, id)
	if err != nil {
		slog.InfoContext(ctx, "%w", err)

		if errors.Is(err, pgx.ErrNoRows) {
			return ErrNotFound
		}

		return err
	}

	return nil
}
