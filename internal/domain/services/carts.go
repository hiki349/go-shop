package services

import (
	"context"
	"go-shop/internal/domain/models"
	"go-shop/internal/storage/repo"

	"github.com/google/uuid"
)

type CartsService struct {
	repo repo.ICartsRepo
}

func NewCartsService(repo repo.ICartsRepo) *CartsService {
	return &CartsService{repo}
}

func (svc CartsService) GetCarts(ctx context.Context) ([]models.Cart, error) {
	carts, err := svc.repo.FindCarts(ctx)
	if err != nil {
		return nil, err
	}

	return carts, nil
}

func (svc CartsService) GetCart(ctx context.Context, id uuid.UUID) (*models.Cart, error) {
	cart, err := svc.repo.FindCartByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (svc CartsService) CreateCart(ctx context.Context, userID uuid.UUID) (*models.Cart, error) {
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
