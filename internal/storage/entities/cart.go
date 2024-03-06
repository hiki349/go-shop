package storage

import "github.com/google/uuid"

type Cart struct {
	ID     uuid.UUID
	UserID string
}
