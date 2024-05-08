package models

import (
	"github.com/google/uuid"

	domain "go-shop/internal/domain/models"
)

type Cart struct {
	ID     uuid.UUID
	Items  []CartItem
	Price  int64
	UserID uuid.UUID
}

type CartItem struct {
	ID       uuid.UUID
	ImageURL string
	Title    string
	Price    int64
	Count    int
}

func (c Cart) ToDomain() *domain.Cart {
	var items []domain.CartItem

	for _, v := range c.Items {
		item := v.ToDomain()
		items = append(items, *item)
	}

	return &domain.Cart{
		ID:     c.ID,
		Items:  items,
		Price:  c.Price,
		UserID: c.UserID,
	}
}

func (ci CartItem) ToDomain() *domain.CartItem {
	return &domain.CartItem{
		ID:       ci.ID,
		ImageURL: ci.ImageURL,
		Title:    ci.Title,
		Price:    ci.Price,
		Count:    ci.Count,
	}
}
