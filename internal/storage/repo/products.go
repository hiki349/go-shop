package repo

import (
	"context"
	"go-shop/internal/domain/models"
	"go-shop/internal/storage/db"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgxutil"
)

type ProductsRepo struct {
	db *db.Postgres
}

type IProductsRepo interface {
	FindAll(ctx context.Context) ([]models.Product, error)
	FindByID(ctx context.Context, id uuid.UUID) (*models.Product, error)
	Create(ctx context.Context, data models.Product) (uuid.UUID, error)
	Update(ctx context.Context, data models.Product) (uuid.UUID, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

func NewProductsRepo(db *db.Postgres) *ProductsRepo {
	return &ProductsRepo{db}
}

func (r ProductsRepo) FindAll(ctx context.Context) ([]models.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	query := "Select id, title, image_url, description, price, created_at, updated_at FROM products;"

	products, err := pgxutil.Select[models.Product](ctx, r.db, query, nil, pgx.RowToStructByPos)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r ProductsRepo) FindByID(ctx context.Context, id uuid.UUID) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	query := `Select
	id, title, image_url, description, price, created_at, updated_at
	FROM products
	WHERE id = $1;`

	product, err := pgxutil.SelectRow[models.Product](ctx, r.db, query, []any{id}, pgx.RowToStructByPos)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductsRepo) Create(ctx context.Context, values models.Product) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	sql := `INSERT INTO products
	(id, title, image_url, price, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6);`

	_, err := pgxutil.ExecRow(
		ctx, r.db, sql,
		values.ID,
		values.Title,
		values.ImageURL,
		values.Price,
		time.Now(),
		nil,
	)
	if err != nil {
		return uuid.UUID{}, err
	}

	return values.ID, nil
}

func (r *ProductsRepo) Update(ctx context.Context, values models.Product) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	sql := `UPDATE products
	SET title = $2, image_url = $3, price = $4, updated_at = $5
	WHERE id = $1;`

	_, err := pgxutil.ExecRow(
		ctx, r.db, sql,
		values.ID,
		values.Title,
		values.ImageURL,
		values.Price,
		values.UpdatedAt,
	)
	if err != nil {
		return values.ID, err
	}

	return values.ID, nil
}

func (r *ProductsRepo) Delete(ctx context.Context, id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	sql := "DELETE FROM products WHERE id = $1;"

	_, err := pgxutil.ExecRow(ctx, r.db, sql, id)
	if err != nil {
		return err
	}

	return nil
}
