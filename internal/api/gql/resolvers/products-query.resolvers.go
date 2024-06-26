package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"
	"errors"
	"go-shop/internal/api/gql/generated"
	"go-shop/internal/api/gql/model"
	"go-shop/internal/storage/repo"
	"log"
	"time"

	"github.com/google/uuid"
)

// GetAll is the resolver for the get_all field.
func (r *productsQueryResolver) GetAll(ctx context.Context, obj *model.ProductsQuery) (model.ProductsFoundResult, error) {
	products, err := r.ProductsService.GetProducts(ctx)
	if err != nil {
		log.Println("%w", err)
		return model.InternalError{Message: "internal error"}, nil
	}

	var res []*model.Product

	for _, v := range products {
		var updatedAt *time.Time

		if !v.UpdatedAt.IsZero() {
			updatedAt = &v.UpdatedAt
		}

		product := &model.Product{
			ID:          v.ID,
			Title:       v.Title,
			ImageURL:    v.ImageURL,
			Description: v.Description,
			Price:       float64(v.Price),
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   updatedAt,
		}

		res = append(res, product)
	}

	return model.ProductsFound{
		Products: res,
	}, nil
}

// GetByID is the resolver for the get_by_id field.
func (r *productsQueryResolver) GetByID(ctx context.Context, obj *model.ProductsQuery, id uuid.UUID) (model.ProductFoundResult, error) {
	var updatedAt *time.Time

	product, err := r.ProductsService.GetProduct(ctx, id)
	if err != nil {
		log.Println("resolver: %w", err)

		if errors.Is(err, repo.ErrNotFound) {
			return model.NotFound{Message: "Product not found"}, nil
		}

		return model.InternalError{Message: "internal error"}, nil
	}

	if !product.CreatedAt.IsZero() {
		updatedAt = &product.UpdatedAt
	}

	resProduct := &model.Product{
		ID:          product.ID,
		Title:       product.Title,
		ImageURL:    product.ImageURL,
		Description: product.Description,
		Price:       float64(product.Price),
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   updatedAt,
	}

	return model.ProductFound{
		Product: resProduct,
	}, nil
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context) (model.ProductsQuery, error) {
	return model.ProductsQuery{}, nil
}

// ProductsQuery returns generated.ProductsQueryResolver implementation.
func (r *Resolver) ProductsQuery() generated.ProductsQueryResolver { return &productsQueryResolver{r} }

type productsQueryResolver struct{ *Resolver }
