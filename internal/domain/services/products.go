package services

import (
	"context"
	"go-shop/internal/domain/models"
	"go-shop/internal/storage/repo"
	"time"

	"github.com/google/uuid"
)

type ProductsService struct {
	repo repo.ProductsRepo
}

func NewProductsService(repo repo.ProductsRepo) *ProductsService {
	return &ProductsService{repo}
}

func (svc ProductsService) GetProducts(ctx context.Context) ([]models.Product, error) {
	products, err := svc.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (svc ProductsService) GetProduct(ctx context.Context, productID uuid.UUID) (*models.Product, error) {
	product, err := svc.repo.FindByID(ctx, productID)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (svc ProductsService) CreateProduct(ctx context.Context, value models.Product) (*models.Product, error) {
	newProduct := value
	newProduct.ID = uuid.New()
	newProduct.CreatedAt = time.Now()

	productID, err := svc.repo.Create(ctx, newProduct)
	if err != nil {
		return nil, err
	}

	product, err := svc.repo.FindByID(ctx, productID)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (svc ProductsService) UpdateProduct(ctx context.Context, id uuid.UUID, value models.Product) (*models.Product, error) {
	updateProduct := value
	updateProduct.ID = id
	updateProduct.CreatedAt = time.Now()

	productID, err := svc.repo.Update(ctx, updateProduct)
	if err != nil {
		return nil, err
	}

	product, err := svc.repo.FindByID(ctx, productID)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (svc ProductsService) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	err := svc.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
