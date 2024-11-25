package types

import "time"

// only cuz i'm adding auth, the task didn't ask for this
type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

type Product struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Seller      string    `json:"seller"`
	Rating      uint8     `json:"rating"`
	CreatedAt   time.Time `json:"createdAt"` // no one asked for this either.
}

type UserStore interface {
	CreateNewUser(User) error
	GetUserByEmail(string) (*User, error)
}

type ProductStore interface {
	GetProducts() ([]*Product, error)
	GetProductByID(int) (*Product, error)
	DeleteProductByID(int) error
	CreateProduct(CreateProductPayload) (*Product, error)
	UpdateProduct(int, UpdateProductPayload) (*Product, error)
}

type CreateProductPayload struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"  validate:"required"`
	Seller      string `json:"seller" validate:"required"`
	Rating      *uint8 `json:"rating,omitempty" validate:"omitempty,gte=1,lte=5"` // this is marked optional in tasks
}

type UpdateProductPayload struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	// Seller      string `json:"seller, omniempty"` // can't update seller ??
	Rating *uint8 `json:"rating,omitempty" validate:"omitempty,gte=1,lte=5"`
}

type NewUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=130"`
}
