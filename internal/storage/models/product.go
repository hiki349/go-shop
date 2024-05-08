package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"

	domain "go-shop/internal/domain/models"
)

type Product struct {
	ID          uuid.UUID
	Title       string
	ImageURL    string
	Description string
	Price       int64
	CreatedAt   time.Time
	UpdatedAt   sql.NullTime
}

func (p Product) ToDomain() *domain.Product {
	return &domain.Product{
		ID:          p.ID,
		Title:       p.Title,
		ImageURL:    p.ImageURL,
		Description: p.Description,
		Price:       p.Price,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   &p.CreatedAt,
	}
}
