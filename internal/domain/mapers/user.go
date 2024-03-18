package mapers

import (
	"go-shop/internal/api/gql/model"
	"go-shop/internal/domain/models"
)

func FromUserToDTO(user models.User) model.UserDto {
	return model.UserDto{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatetAt,
	}
}

func FromReqToUser(req model.UserReq) models.User {
	return models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
}
