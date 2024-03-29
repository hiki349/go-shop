package services

import (
	"context"
	"go-shop/internal/api/gql/model"
	"go-shop/internal/domain/mapers"
	"go-shop/internal/storage/repo"
	"time"

	"github.com/google/uuid"
)

type ProductsService struct {
	repo repo.IProductsRepo
}

func NewProductsService(repo repo.IProductsRepo) *ProductsService {
	return &ProductsService{repo}
}

func (svc ProductsService) GetProducts(ctx context.Context) ([]*model.Product, error) {
	products, err := svc.repo.FindProducts(ctx)
	if err != nil {
		return nil, err
	}

	productsDTO := mapers.FromProductsToDTO(products)

	return productsDTO, nil
}

func (svc ProductsService) GetProduct(ctx context.Context, productID uuid.UUID) (*model.Product, error) {
	product, err := svc.repo.FindProductByID(ctx, productID)
	if err != nil {
		return nil, err
	}

	productDTO := mapers.FromProductToDTO(*product)

	return &productDTO, nil
}

func (svc ProductsService) CreateProduct(ctx context.Context, value model.ProductReq) (*model.Product, error) {
	newProduct := mapers.FromReqToProduct(value)
	newProduct.ID = uuid.New()
	newProduct.CreatedAt = time.Now()

	productID, err := svc.repo.CreateProduct(ctx, newProduct)
	if err != nil {
		return nil, err
	}

	product, err := svc.repo.FindProductByID(ctx, productID)
	if err != nil {
		return nil, err
	}

	productDTO := mapers.FromProductToDTO(*product)

	return &productDTO, nil
}

func (svc ProductsService) UpdateProduct(ctx context.Context, id uuid.UUID, value model.ProductReq) (*model.Product, error) {
	updateProduct := mapers.FromReqToProduct(value)
	updateProduct.ID = id
	updateProduct.CreatedAt = time.Now()

	productID, err := svc.repo.UpdateProduct(ctx, updateProduct)
	if err != nil {
		return nil, err
	}

	product, err := svc.repo.FindProductByID(ctx, productID)
	if err != nil {
		return nil, err
	}

	productDTO := mapers.FromProductToDTO(*product)

	return &productDTO, nil
}

func (svc ProductsService) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	err := svc.repo.DeleteProduct(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
