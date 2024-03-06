package models

import "github.com/google/uuid"

type Cart struct {
	ID    uuid.UUID
	Items []CartItem
	Price int64
}

type CartItem struct {
	ID       uuid.UUID
	ImageURL string
	Title    string
	Price    int64
	Count    int
}
