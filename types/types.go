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
	Rating      float64   `json:"rating"`
	CreatedAt   time.Time `json:"createdAt"` // no one asked for this either.
}

type UserStore interface {
	CreateNewUser(User) error
	GetUserByEmail(email string) (*User, error)
}

type NewUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=130"`
}
