package mapers

import (
	"go-shop/internal/api/gql/model"
	"go-shop/internal/domain/models"
)

func ToCartDTO(cart models.Cart) model.Cart {
	var cartItemsDTO []*model.CartItem
	for _, item := range cart.Items {
		cartItemsDTO = append(
			cartItemsDTO,
			toCartItemDTO(item),
		)
	}

	return model.Cart{
		ID:     cart.ID.String(),
		Items:  cartItemsDTO,
		Price:  float64(cart.Price / 100),
		UserID: cart.ID.String(),
	}
}

func toCartItemDTO(cartItem models.CartItem) *model.CartItem {
	return &model.CartItem{
		ID:       cartItem.ID.String(),
		Title:    cartItem.Title,
		ImageURL: cartItem.ImageURL,
		Price:    float64(cartItem.Price / 100),
		Count:    cartItem.Count,
	}
}
