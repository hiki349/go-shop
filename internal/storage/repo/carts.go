package repo

import (
	"context"
	"errors"
	"log/slog"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgxutil"

	"go-shop/internal/domain/models"
	"go-shop/internal/storage/db"
)

var (
	ErrCartNotFound = errors.New("not found cart")
)

type PostgresCartsRepo struct {
	db *db.Postgres
}

type CartsRepo interface {
	FindAll(ctx context.Context) ([]models.Cart, error)
	FindByID(ctx context.Context, id uuid.UUID) (*models.Cart, error)
	Create(ctx context.Context, cartID, userID uuid.UUID) (uuid.UUID, error)
}

func NewPostgresCartsRepo(db *db.Postgres) *PostgresCartsRepo {
	return &PostgresCartsRepo{db}
}

func (r PostgresCartsRepo) FindAll(ctx context.Context) ([]models.Cart, error) {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	query := `SELECT carts.id, carts.price, carts.user_id,
	cart_items.id, cart_items.image_url, cart_items.title, cart_items.price, cart_items.count
	FROM carts
	INNER JOIN cart_items
	ON cart_items.cart_id = carts.id`

	products, err := pgxutil.Select(ctx, r.db, query, nil, pgx.RowToStructByPos[models.Cart])
	if err != nil {
		slog.InfoContext(ctx, "%w", err)

		return nil, err
	}

	return products, nil
}

func (r PostgresCartsRepo) FindByID(ctx context.Context, id uuid.UUID) (*models.Cart, error) {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	query := `SELECT carts.id, carts.price, carts.user_id,
	cart_items.id, cart_items.image_url, cart_items.title, cart_items.price, cart_items.count
	FROM carts
	WHERE id = $1
	INNER JOIN cart_items
	ON cart_items.cart_id = carts.id`

	product, err := pgxutil.SelectRow[models.Cart](ctx, r.db, query, []any{id}, pgx.RowToStructByPos)
	if err != nil {
		slog.InfoContext(ctx, "%w", err)

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrCartNotFound
		}

		return nil, err
	}

	return &product, nil
}

func (r *PostgresCartsRepo) Create(ctx context.Context, cartID, userID uuid.UUID) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	sql := "INSERT INTO carts (id, user_id) VALUES ($1, $2);"

	_, err := pgxutil.ExecRow(
		ctx, r.db, sql,
		cartID,
		userID,
	)
	if err != nil {
		slog.InfoContext(ctx, "%w", err)

		if errors.Is(err, pgx.ErrNoRows) {
			return uuid.UUID{}, ErrCartNotFound
		}

		return uuid.UUID{}, err
	}

	return cartID, nil
}
