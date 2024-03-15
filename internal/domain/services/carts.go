package services

import (
	"context"
	"go-shop/internal/domain/models"

	"github.com/google/uuid"
)

func (svc Services) GetCarts(ctx context.Context) ([]models.Cart, error) {
	carts, err := svc.repo.FindCarts(ctx)
	if err != nil {
		return nil, err
	}

	return carts, nil
}

func (svc Services) GetCart(ctx context.Context, id uuid.UUID) (*models.Cart, error) {
	cart, err := svc.repo.FindCartByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (svc *Services) CreateCart(ctx context.Context, userID uuid.UUID) (*models.Cart, error) {
	cartID := uuid.New()
	cartIDFromDB, err := svc.repo.CreateCart(ctx, cartID, userID)
	if err != nil {
		return nil, err
	}

	cart, err := svc.repo.FindCartByID(ctx, cartIDFromDB)
	if err != nil {
		return nil, err
	}

	return cart, nil
}
