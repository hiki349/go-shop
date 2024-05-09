package services

import (
	"context"
	"time"

	"github.com/google/uuid"

	"go-shop/graph/model"
	"go-shop/internal/domain/models"
	"go-shop/internal/storage/repo"
)

type UsersService struct {
	repo repo.UsersRepo
}

type IUsersService interface {
	GetUsers(ctx context.Context) ([]models.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	CreateUser(ctx context.Context, value model.NewUser) (*models.User, error)
	UpdateUser(ctx context.Context, id uuid.UUID, value models.User) (*models.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

func NewUsersService(repo repo.UsersRepo) *UsersService {
	return &UsersService{repo}
}

func (svc UsersService) GetUsers(ctx context.Context) ([]models.User, error) {
	users, err := svc.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (svc UsersService) GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	user, err := svc.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (svc *UsersService) CreateUser(ctx context.Context, value model.NewUser) (*models.User, error) {
	var updatedAt time.Time

	newUser := &models.User{
		ID:        uuid.New(),
		Username:  value.Username,
		Email:     value.Email,
		Password:  value.Password,
		CreatedAt: time.Now(),
		UpdatedAt: updatedAt,
	}

	userID, err := svc.repo.Create(ctx, newUser)
	if err != nil {
		return nil, err
	}

	user, err := svc.repo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (svc *UsersService) UpdateUser(ctx context.Context, id uuid.UUID, value model.NewUser) (*models.User, error) {
	updateUser := &models.User{
		ID: id,
		Username: value.Username,
		Email: value.Email,
		Password: value.Password,
		UpdatedAt: time.Now(),
	}

	userID, err := svc.repo.Update(ctx, updateUser)
	if err != nil {
		return nil, err
	}

	user, err := svc.repo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (svc *UsersService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	err := svc.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
