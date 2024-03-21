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
	db *db.DB
}

type ICartsRepo interface {
	FindCarts(ctx context.Context) ([]models.Cart, error)
	FindCartByID(ctx context.Context, id uuid.UUID) (*models.Cart, error)
	CreateCart(ctx context.Context, cartID, userID uuid.UUID) (uuid.UUID, error)
}

func NewCartsRepo(db *db.DB) *CartsRepo {
	return &CartsRepo{db}
}

func (r CartsRepo) FindCarts(ctx context.Context) ([]models.Cart, error) {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	query := `Select carts.id, carts.price, carts.user_id,
	cart_items.id, cart_items.image_url, cart_items.title, cart_items.price, cart_items.count
	FROM carts
	INNER JOIN cart_items
	ON cart_items.cart_id = carts.id`

	products, err := pgxutil.Select(ctx, r.db.Postgres, query, nil, pgx.RowToStructByPos[models.Cart])
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r CartsRepo) FindCartByID(ctx context.Context, id uuid.UUID) (*models.Cart, error) {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	query := `Select carts.id, carts.price, carts.user_id,
	cart_items.id, cart_items.image_url, cart_items.title, cart_items.price, cart_items.count
	FROM carts
	WHERE id = $1
	INNER JOIN cart_items
	ON cart_items.cart_id = carts.id`

	product, err := pgxutil.SelectRow(ctx, r.db.Postgres, query, []any{id}, pgx.RowToStructByPos[models.Cart])
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *CartsRepo) CreateCart(ctx context.Context, cartID, userID uuid.UUID) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	sql := "INSERT INTO carts (id, user_id) VALUES ($1, $2);"

	_, err := pgxutil.ExecRow(
		ctx, r.db.Postgres, sql,
		cartID,
		userID,
	)
	if err != nil {
		return uuid.UUID{}, err
	}

	return cartID, nil
}
