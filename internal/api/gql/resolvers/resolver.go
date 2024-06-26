package resolvers

import "go-shop/internal/domain/services"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ProductsService *services.ProductsService
	UsersService    *services.UsersService
	CartsService    *services.CartsService
}
