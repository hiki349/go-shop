package mapers

import (
	"go-shop/internal/api/gql/model"
	"go-shop/internal/domain/models"
)

func FromProductToDTO(product models.Product) model.Product {
	return model.Product{
		ID:          product.ID,
		Title:       product.Title,
		ImageURL:    product.ImageURL,
		Description: product.Description,
		Price:       float64(product.Price / 100),
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

func FromProductsToDTO(products []models.Product) []*model.Product {
	var productsDTO []*model.Product
	for _, item := range products {
		itemDTO := FromProductToDTO(item)
		productsDTO = append(productsDTO, &itemDTO)
	}

	return productsDTO
}

func FromReqToProduct(req model.Product) models.Product {
	return models.Product{
		Title:       req.Title,
		ImageURL:    req.ImageURL,
		Description: req.Description,
		Price:       int64(req.Price * 100),
	}
}
