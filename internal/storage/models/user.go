package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"

	domain "go-shop/internal/domain/models"
)

type User struct {
	ID        uuid.UUID
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

func (u User) ToDomain() *domain.User {
	var updatedAt time.Time

	if u.UpdatedAt.Valid {
		updatedAt = u.UpdatedAt.Time
	}

	return &domain.User{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: updatedAt,
	}
}
