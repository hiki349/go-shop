package repo

import (
	"context"
	"go-shop/internal/domain/models"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgxutil"
)

func (r Repo) FindProducts(ctx context.Context) ([]models.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	query := "Select id, title, image_url, description, price, created_at, updated_at FROM products;"

	products, err := pgxutil.Select[models.Product](ctx, r.db.Postgres, query, nil, pgx.RowToStructByPos)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r Repo) FindProductByID(ctx context.Context, id uuid.UUID) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	query := `Select
	id, title, image_url, description, price, created_at, updated_at
	FROM products
	WHERE id = $1;`

	product, err := pgxutil.SelectRow[models.Product](ctx, r.db.Postgres, query, []any{id}, pgx.RowToStructByPos)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *Repo) CreateProduct(ctx context.Context, values models.Product) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	sql := `INSERT INTO products
	(id, title, image_url, price, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6);`

	productID := uuid.New()
	_, err := pgxutil.ExecRow(
		ctx, r.db.Postgres, sql,
		productID,
		values.Title,
		values.ImageURL,
		values.Price,
		time.Now(),
		nil,
	)
	if err != nil {
		return uuid.UUID{}, err
	}

	return productID, nil
}

func (r *Repo) UpdateProduct(ctx context.Context, values models.Product) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	sql := `UPDATE products
	SET title = $2, image_url = $3, price = $4, updated_at = $5
	WHERE id = $1;`

	_, err := pgxutil.ExecRow(
		ctx, r.db.Postgres, sql,
		values.ID,
		values.Title,
		values.ImageURL,
		values.Price,
		time.Now(),
	)
	if err != nil {
		return values.ID, err
	}

	return values.ID, nil
}

func (r *Repo) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	sql := "DELETE FROM products WHERE id = $1;"

	_, err := pgxutil.ExecRow(ctx, r.db.Postgres, sql, id)
	if err != nil {
		return err
	}

	return nil
}
