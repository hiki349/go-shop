package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"

	"go-shop/graph/model"
)

// User is the resolver for the user field.
func (r *mutationResolver) User(ctx context.Context) (model.UserMutation, error) {
	return r.Mutation().User(ctx)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) (model.UsersQuery, error) {
	return r.Query().Users(ctx)
}

// Create is the resolver for the create field.
func (r *userMutationResolver) Create(ctx context.Context, obj *model.UserMutation, input model.NewUser) (*model.User, error) {
	var updatedAt *time.Time

	user, err := r.UsersService.CreateUser(ctx, input)
	if err != nil {
		return nil, err
	}

	if !user.UpdatedAt.IsZero() {
		updatedAt = &user.UpdatedAt
	}

	return &model.User{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatetAt: updatedAt,
	}, nil
}

// Update is the resolver for the update field.
func (r *userMutationResolver) Update(ctx context.Context, obj *model.UserMutation, id uuid.UUID, input model.NewUser) (*model.User, error) {
	var updatedAt *time.Time

	user, err := r.UsersService.UpdateUser(ctx, id, input)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if !user.UpdatedAt.IsZero() {
		updatedAt = &user.UpdatedAt
	}

	return &model.User{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatetAt: updatedAt,
	}, nil
}

// Delete is the resolver for the delete field.
func (r *userMutationResolver) Delete(ctx context.Context, obj *model.UserMutation, id uuid.UUID) (bool, error) {
	err := r.UsersService.DeleteUser(ctx, id)
	if err != nil {
		return false, err
	}

	return true, nil
}

// GetByID is the resolver for the get_by_id field.
func (r *usersQueryResolver) GetByID(ctx context.Context, obj *model.UsersQuery, id uuid.UUID) (*model.User, error) {
	var updatedAt *time.Time

	user, err := r.UsersService.GetUserByID(ctx, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if !user.UpdatedAt.IsZero() {
		updatedAt = &user.UpdatedAt
	}

	return &model.User{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatetAt: updatedAt,
	}, nil
}

// GetAll is the resolver for the get_all field.
func (r *usersQueryResolver) GetAll(ctx context.Context, obj *model.UsersQuery) ([]*model.User, error) {
	users, err := r.UsersService.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	var res []*model.User
	for _, v := range users {
		var updatedAt *time.Time

		if !v.UpdatedAt.IsZero() {
			updatedAt = &v.UpdatedAt
		}

		user := &model.User{
			ID:        v.ID,
			Username:  v.Username,
			Email:     v.Email,
			Password:  v.Password,
			CreatedAt: v.CreatedAt,
			UpdatetAt: updatedAt,
		}

		res = append(res, user)
	}

	return res, nil
}

// UserMutation returns UserMutationResolver implementation.
func (r *Resolver) UserMutation() UserMutationResolver { return &userMutationResolver{r} }

// UsersQuery returns UsersQueryResolver implementation.
func (r *Resolver) UsersQuery() UsersQueryResolver { return &usersQueryResolver{r} }

type userMutationResolver struct{ *Resolver }
type usersQueryResolver struct{ *Resolver }
