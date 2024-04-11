package repo

import (
	"context"
	"go-shop/internal/domain/models"
	"go-shop/internal/storage/db"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgxutil"
)

type CartsRepo struct {
	db *db.Postgres
}

type ICartsRepo interface {
	FindAll(ctx context.Context) ([]models.Cart, error)
	FindByID(ctx context.Context, id uuid.UUID) (*models.Cart, error)
	Create(ctx context.Context, cartID, userID uuid.UUID) (uuid.UUID, error)
}

func NewCartsRepo(db *db.Postgres) *CartsRepo {
	return &CartsRepo{db}
}

func (r CartsRepo) FindAll(ctx context.Context) ([]models.Cart, error) {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	query := `Select carts.id, carts.price, carts.user_id,
	cart_items.id, cart_items.image_url, cart_items.title, cart_items.price, cart_items.count
	FROM carts
	INNER JOIN cart_items
	ON cart_items.cart_id = carts.id`

	products, err := pgxutil.Select(ctx, r.db, query, nil, pgx.RowToStructByPos[models.Cart])
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r CartsRepo) FindByID(ctx context.Context, id uuid.UUID) (*models.Cart, error) {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	query := `Select carts.id, carts.price, carts.user_id,
	cart_items.id, cart_items.image_url, cart_items.title, cart_items.price, cart_items.count
	FROM carts
	WHERE id = $1
	INNER JOIN cart_items
	ON cart_items.cart_id = carts.id`

	product, err := pgxutil.SelectRow(ctx, r.db, query, []any{id}, pgx.RowToStructByPos[models.Cart])
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *CartsRepo) Create(ctx context.Context, cartID, userID uuid.UUID) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	sql := "INSERT INTO carts (id, user_id) VALUES ($1, $2);"

	_, err := pgxutil.ExecRow(
		ctx, r.db, sql,
		cartID,
		userID,
	)
	if err != nil {
		return uuid.UUID{}, err
	}

	return cartID, nil
}
