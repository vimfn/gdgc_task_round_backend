package user

import (
	"database/sql"

	"vitshop.vimfn.in/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateNewUser(u types.User) error {
	_, err := s.db.Exec("INSERT INTO users (email, password) VALUES (?, ?)", u.Email, u.Password)
	if err != nil {
		return err
	}

	return nil
}
