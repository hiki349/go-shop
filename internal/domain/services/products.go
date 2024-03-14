package services

import (
	"context"
	"go-shop/internal/domain/models"

	"github.com/google/uuid"
)

func (svc Services) GetProducts(ctx context.Context) ([]models.Product, error) {
	products, err := svc.repo.FindProducts(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (svc Services) GetProduct(ctx context.Context, productID uuid.UUID) (*models.Product, error) {
	product, err := svc.repo.FindProductByID(ctx, productID)
	if err != nil {
		return nil, err
	}

	return product, nil
}
