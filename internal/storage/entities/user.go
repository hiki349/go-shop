package storage

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatetAt time.Time
}
