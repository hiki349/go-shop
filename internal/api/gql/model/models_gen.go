// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"

	"github.com/google/uuid"
)

type ProblemInterface interface {
	IsProblemInterface()
	GetMessage() string
}

type ProductResult interface {
	IsProductResult()
}

type ProductsCreateResult interface {
	IsProductsCreateResult()
}

type ProductsDeleteResult interface {
	IsProductsDeleteResult()
}

type ProductsResult interface {
	IsProductsResult()
}

type ProductsUpdateResult interface {
	IsProductsUpdateResult()
}

type Cart struct {
	ID     uuid.UUID   `json:"id"`
	Items  []*CartItem `json:"items"`
	Price  float64     `json:"price"`
	UserID string      `json:"user_id"`
}

type CartItem struct {
	ID       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
	ImageURL string    `json:"imageUrl"`
	Price    float64   `json:"price"`
	Count    int       `json:"count"`
}

type InternalErrorProblem struct {
	Message string `json:"message"`
}

func (InternalErrorProblem) IsProblemInterface()     {}
func (this InternalErrorProblem) GetMessage() string { return this.Message }

func (InternalErrorProblem) IsProductsCreateResult() {}

func (InternalErrorProblem) IsProductsUpdateResult() {}

func (InternalErrorProblem) IsProductsDeleteResult() {}

func (InternalErrorProblem) IsProductsResult() {}

func (InternalErrorProblem) IsProductResult() {}

type Mutation struct {
}

type Product struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	ImageURL    string    `json:"imageUrl"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type ProductCreateOk struct {
	Product *Product `json:"product"`
}

func (ProductCreateOk) IsProductsCreateResult() {}

type ProductDeleteOk struct {
	ProductID uuid.UUID `json:"productId"`
}

func (ProductDeleteOk) IsProductsDeleteResult() {}

type ProductInput struct {
	Title       string  `json:"title"`
	ImageURL    string  `json:"imageUrl"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type ProductNotFound struct {
	Message string `json:"message"`
}

func (ProductNotFound) IsProblemInterface()     {}
func (this ProductNotFound) GetMessage() string { return this.Message }

func (ProductNotFound) IsProductsUpdateResult() {}

func (ProductNotFound) IsProductsDeleteResult() {}

func (ProductNotFound) IsProductResult() {}

type ProductOk struct {
	Product *Product `json:"product"`
}

func (ProductOk) IsProductResult() {}

type ProductUpdateOk struct {
	Product *Product `json:"product"`
}

func (ProductUpdateOk) IsProductsUpdateResult() {}

type ProductsMutation struct {
	Create ProductsCreateResult `json:"create"`
	Update ProductsUpdateResult `json:"update"`
	Delete ProductsDeleteResult `json:"delete"`
}

type ProductsOk struct {
	Products []*Product `json:"products"`
}

func (ProductsOk) IsProductsResult() {}

type ProductsQuery struct {
	FindAll ProductsResult `json:"findAll"`
	Find    *ProductOk     `json:"find"`
}

type Query struct {
}

type UserDto struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
