package repo

import (
	"context"
	"go-shop/internal/domain/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgxutil"
)

func (r Repo) FindCarts(ctx context.Context) ([]models.Cart, error) {
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

func (r Repo) FindCartByID(ctx context.Context, id string) (*models.Cart, error) {
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

func (r *Repo) CreateCart(ctx context.Context, userID uuid.UUID) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	sql := "INSERT INTO carts (id, user_id) VALUES ($1, $2);"

	cartID := uuid.New()
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
