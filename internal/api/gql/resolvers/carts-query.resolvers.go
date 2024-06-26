package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"
	"go-shop/internal/api/gql/generated"
	"go-shop/internal/api/gql/model"

	"github.com/google/uuid"
)

// GetAll is the resolver for the get_all field.
func (r *cartsQueryResolver) GetAll(ctx context.Context, obj *model.CartsQuery) (model.CartsFoundResult, error) {
	var res []*model.Cart

	carts, err := r.CartsService.GetCarts(ctx)
	if err != nil {
		return model.InternalError{Message: "internal service error"}, nil
	}

	for _, v := range carts {
		product := model.Cart{
			ID:     v.ID,
			UserID: v.UserID,
		}

		res = append(res, &product)
	}

	return model.CartsFound{
		Carts: res,
	}, nil
}

// GetByID is the resolver for the get_by_id field.
func (r *cartsQueryResolver) GetByID(ctx context.Context, obj *model.CartsQuery, id uuid.UUID) (model.CartFoundResult, error) {
	cart, err := r.CartsService.GetCart(ctx, id)
	if err != nil {
		return model.InternalError{Message: "internal service error"}, nil
	}

	res := &model.Cart{
		ID:     cart.ID,
		UserID: cart.UserID,
	}

	return model.CartFound{
		Cart: res,
	}, nil
}

// Carts is the resolver for the carts field.
func (r *queryResolver) Carts(ctx context.Context) (model.CartsQuery, error) {
	return model.CartsQuery{}, nil
}

// CartsQuery returns generated.CartsQueryResolver implementation.
func (r *Resolver) CartsQuery() generated.CartsQueryResolver { return &cartsQueryResolver{r} }

type cartsQueryResolver struct{ *Resolver }
