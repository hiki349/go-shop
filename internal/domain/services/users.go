package services

import (
	"context"
	"go-shop/internal/api/gql/model"
	"go-shop/internal/domain/mapers"
	"go-shop/internal/domain/models"
	"go-shop/internal/storage/repo"
	"time"

	"github.com/google/uuid"
)

type UsersService struct {
	repo repo.IUsersRepo
}

func NewUsersService(repo repo.IUsersRepo) *UsersService {
	return &UsersService{repo}
}

func (svc UsersService) GetUsers(ctx context.Context) ([]models.User, error) {
	users, err := svc.repo.FindUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (svc UsersService) GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	user, err := svc.repo.FindUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (svc *UsersService) CreateUser(ctx context.Context, value model.UserReq) (*models.User, error) {
	newUser := mapers.FromReqToUser(value)
	newUser.ID = uuid.New()
	newUser.CreatedAt = time.Now()

	userID, err := svc.repo.CreateUser(ctx, newUser)
	if err != nil {
		return nil, err
	}

	user, err := svc.repo.FindUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (svc *UsersService) UpdateUser(ctx context.Context, id uuid.UUID, value model.UserReq) (*models.User, error) {
	updateUser := mapers.FromReqToUser(value)
	updateUser.ID = uuid.New()
	updateUser.UpdatetAt = time.Now()

	userID, err := svc.repo.UpdateUser(ctx, updateUser)
	if err != nil {
		return nil, err
	}

	user, err := svc.repo.FindUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (svc *UsersService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	err := svc.repo.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
