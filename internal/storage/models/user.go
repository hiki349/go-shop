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
	UpdatetAt sql.NullTime
}

func (u User) ToDomain() *domain.User {
	return &domain.User{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatetAt: &u.UpdatetAt.Time,
	}
}
