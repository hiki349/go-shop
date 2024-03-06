package storage

import "github.com/google/uuid"

type CartItem struct {
	ID        uuid.UUID
	ProductID string
	CartID    string
}
