package services

import (
	"context"
	"go-shop/internal/domain/models"
	"time"

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

func (svc *Services) CreateProduct(ctx context.Context, value models.Product) (*models.Product, error) {
	newProduct := models.Product{
		ID:          uuid.New(),
		Title:       value.Title,
		ImageURL:    value.ImageURL,
		Description: value.Description,
		Price:       value.Price,
		CreatedAt:   time.Now(),
	}

	productID, err := svc.repo.CreateProduct(ctx, newProduct)
	if err != nil {
		return nil, err
	}

	return svc.repo.FindProductByID(ctx, productID)
}

func (svc *Services) UpdateProduct(ctx context.Context, id uuid.UUID, value models.Product) (*models.Product, error) {
	updateProduct := models.Product{
		ID:          id,
		Title:       value.Title,
		ImageURL:    value.ImageURL,
		Description: value.Description,
		Price:       value.Price,
		UpdatedAt:   time.Now(),
	}

	productID, err := svc.repo.UpdateProduct(ctx, updateProduct)
	if err != nil {
		return nil, err
	}

	return svc.repo.FindProductByID(ctx, productID)
}

func (svc *Services) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	err := svc.repo.DeleteProduct(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
