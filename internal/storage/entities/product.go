package storage

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID
	Title       string
	ImageURL    string
	Description string
	Price       int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
