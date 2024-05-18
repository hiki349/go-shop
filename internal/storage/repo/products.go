package repo

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgxutil"

	domain "go-shop/internal/domain/models"
	"go-shop/internal/storage/db"
	"go-shop/internal/storage/models"
)

var (
	ErrNotFound = errors.New("not found")
)

type PostgresProductsRepo struct {
	db *db.Postgres
}

type ProductsRepo interface {
	FindAll(ctx context.Context) ([]domain.Product, error)
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Product, error)
	Create(ctx context.Context, data domain.Product) (uuid.UUID, error)
	Update(ctx context.Context, data domain.Product) (uuid.UUID, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

func NewPostgresProductsRepo(db *db.Postgres) *PostgresProductsRepo {
	return &PostgresProductsRepo{db}
}

func (r PostgresProductsRepo) FindAll(ctx context.Context) ([]domain.Product, error) {
	var domainProducts []domain.Product

	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	query := "SELECT id, title, image_url, description, price, created_at, updated_at FROM products;"

	products, err := pgxutil.Select[models.Product](ctx, r.db, query, nil, pgx.RowToStructByPos)
	if err != nil {
		return nil, err
	}

	for _, v := range products {
		product := v.ToDomain()
		domainProducts = append(domainProducts, *product)
	}

	return domainProducts, nil
}

func (r PostgresProductsRepo) FindByID(ctx context.Context, id uuid.UUID) (*domain.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	query := `SELECT
	id, title, image_url, description, price, created_at, updated_at
	FROM products
	WHERE id = $1;`

	product, err := pgxutil.SelectRow[models.Product](ctx, r.db, query, []any{id}, pgx.RowToStructByPos)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return product.ToDomain(), nil
}

func (r *PostgresProductsRepo) Create(ctx context.Context, values domain.Product) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	sql := `INSERT INTO products
	(id, title, image_url, description, price, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7);`

	_, err := pgxutil.ExecRow(
		ctx, r.db, sql,
		values.ID,
		values.Title,
		values.ImageURL,
		values.Description,
		values.Price,
		time.Now(),
		nil,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return uuid.UUID{}, ErrNotFound
		}

		return uuid.UUID{}, err
	}

	return values.ID, nil
}

func (r *PostgresProductsRepo) Update(ctx context.Context, values domain.Product) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	sql := `UPDATE products
	SET title = $2, image_url = $3, description = $4, price = $5, updated_at = $6
	WHERE id = $1;`

	_, err := pgxutil.ExecRow(
		ctx, r.db, sql,
		values.ID,
		values.Title,
		values.ImageURL,
		values.Description,
		values.Price,
		values.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return uuid.UUID{}, ErrNotFound
		}

		return uuid.UUID{}, err
	}

	return values.ID, nil
}

func (r *PostgresProductsRepo) Delete(ctx context.Context, id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, maxTimeToDoDbOperation)
	defer cancel()

	sql := "DELETE FROM products WHERE id = $1;"

	_, err := pgxutil.ExecRow(ctx, r.db, sql, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrNotFound
		}

		return err
	}

	return nil
}
