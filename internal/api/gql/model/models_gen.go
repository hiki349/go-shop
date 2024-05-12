// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"

	"github.com/google/uuid"
)

type ErrorInterface interface {
	IsErrorInterface()
	GetMessage() string
}

type ProductCreateResult interface {
	IsProductCreateResult()
}

type ProductDeleteResult interface {
	IsProductDeleteResult()
}

type ProductFoundResult interface {
	IsProductFoundResult()
}

type ProductUpdateResult interface {
	IsProductUpdateResult()
}

type ProductsFoundResult interface {
	IsProductsFoundResult()
}

type UserCreateResult interface {
	IsUserCreateResult()
}

type UserDeleteResult interface {
	IsUserDeleteResult()
}

type UserFoundResult interface {
	IsUserFoundResult()
}

type UserUpdateResult interface {
	IsUserUpdateResult()
}

type UsersFoundResult interface {
	IsUsersFoundResult()
}

type InternalError struct {
	Message string `json:"message"`
}

func (InternalError) IsProductCreateResult() {}

func (InternalError) IsProductUpdateResult() {}

func (InternalError) IsProductDeleteResult() {}

func (InternalError) IsProductFoundResult() {}

func (InternalError) IsProductsFoundResult() {}

func (InternalError) IsErrorInterface()       {}
func (this InternalError) GetMessage() string { return this.Message }

func (InternalError) IsUserCreateResult() {}

func (InternalError) IsUserUpdateResult() {}

func (InternalError) IsUserDeleteResult() {}

func (InternalError) IsUserFoundResult() {}

func (InternalError) IsUsersFoundResult() {}

type Mutation struct {
}

type NewProduct struct {
	Title       string  `json:"title"`
	ImageURL    string  `json:"image_url"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type NewUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NotFound struct {
	Message string `json:"message"`
}

func (NotFound) IsProductUpdateResult() {}

func (NotFound) IsProductDeleteResult() {}

func (NotFound) IsProductFoundResult() {}

func (NotFound) IsErrorInterface()       {}
func (this NotFound) GetMessage() string { return this.Message }

func (NotFound) IsUserUpdateResult() {}

func (NotFound) IsUserDeleteResult() {}

func (NotFound) IsUserFoundResult() {}

type Product struct {
	ID          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	ImageURL    string     `json:"image_url"`
	Description string     `json:"description"`
	Price       float64    `json:"price"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

type ProductCreate struct {
	Product *Product `json:"product"`
}

func (ProductCreate) IsProductCreateResult() {}

type ProductDelete struct {
	IsDelete bool `json:"is_delete"`
}

func (ProductDelete) IsProductDeleteResult() {}

type ProductFound struct {
	Product *Product `json:"product"`
}

func (ProductFound) IsProductFoundResult() {}

type ProductMutation struct {
	Create ProductCreateResult `json:"create"`
	Update ProductUpdateResult `json:"update"`
	Delete ProductDeleteResult `json:"delete"`
}

type ProductUpdate struct {
	Product *Product `json:"product"`
}

func (ProductUpdate) IsProductUpdateResult() {}

type ProductsFound struct {
	Products []*Product `json:"products"`
}

func (ProductsFound) IsProductsFoundResult() {}

type ProductsQuery struct {
	GetAll  ProductsFoundResult `json:"get_all"`
	GetByID ProductFoundResult  `json:"get_by_id"`
}

type Query struct {
}

type User struct {
	ID        uuid.UUID  `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatetAt *time.Time `json:"updatet_at,omitempty"`
}

type UserCreate struct {
	User *User `json:"user"`
}

func (UserCreate) IsUserCreateResult() {}

type UserDelete struct {
	IsDelete bool `json:"is_delete"`
}

func (UserDelete) IsUserDeleteResult() {}

type UserFound struct {
	User *User `json:"user"`
}

func (UserFound) IsUserFoundResult() {}

type UserMutation struct {
	Create UserCreateResult `json:"create"`
	Update UserUpdateResult `json:"update"`
	Delete UserDeleteResult `json:"delete"`
}

type UserUpdate struct {
	User *User `json:"user"`
}

func (UserUpdate) IsUserUpdateResult() {}

type UsersFound struct {
	Users []*User `json:"users"`
}

func (UsersFound) IsUsersFoundResult() {}

type UsersQuery struct {
	GetByID UserFoundResult  `json:"get_by_id"`
	GetAll  UsersFoundResult `json:"get_all"`
}
