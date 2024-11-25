package types

import "time"

type Product struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Seller      string    `json:"seller"`
	Rating      float64   `json:"rating"`
	CreatedAt   time.Time `json:"createdAt"` // no one asked for this either.
}

type ProductStore interface {
	GetProducts() ([]*Product, error)
	GetProductByID(int) (*Product, error)
	DeleteProductByID(int) error
	CreateProduct(CreateProductPayload) (*Product, error)
	UpdateProduct(int, UpdateProductPayload) (*Product, error)
}

type CreateProductPayload struct {
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description"  validate:"required"`
	Seller      string   `json:"seller" validate:"required"`
	Rating      *float64 `json:"rating,omitempty" validate:"omitempty,gte=1,lte=5"` // this is marked optional in tasks
}

type UpdateProductPayload struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	// Seller      string `json:"seller, omniempty"` // can't update seller ??
	Rating *uint8 `json:"rating,omitempty" validate:"omitempty,gte=1,lte=5"`
}
