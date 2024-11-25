package product

import (
	"database/sql"
	"fmt"
	"strings"

	"vitshop.vimfn.in/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProductByID(productID int) (*types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products WHERE id = ?", productID)
	if err != nil {
		return nil, err
	}

	p := new(types.Product)
	for rows.Next() {
		p, err = scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
	}

	return p, nil
}

func scanRowsIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)

	err := rows.Scan(
		&product.ID,
		&product.Title,
		&product.Description,
		&product.Seller,
		&product.Rating,
		&product.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *Store) CreateProduct(product types.CreateProductPayload) (*types.Product, error) {
	rating := float64(0)
	if product.Rating != nil {
		rating = *product.Rating
	}

	_, err := s.db.Exec("INSERT INTO products (title, description, seller, rating) VALUES (?, ?, ?, ?)",
		product.Title, product.Description, product.Seller, rating)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *Store) UpdateProduct(productID int, payload types.UpdateProductPayload) (*types.Product, error) {
	currentProduct, err := s.GetProductByID(productID)
	if err != nil {
		return nil, err
	}

	// build the update query dynamically based on which fields are provided
	updates := []string{}
	args := []interface{}{}

	if payload.Title != nil {
		updates = append(updates, "title = ?")
		args = append(args, *payload.Title)
	}
	if payload.Description != nil {
		updates = append(updates, "description = ?")
		args = append(args, *payload.Description)
	}
	if payload.Rating != nil {
		updates = append(updates, "rating = ?")
		args = append(args, *payload.Rating)
	}

	if len(updates) == 0 {
		return currentProduct, nil
	}

	args = append(args, productID)

	query := fmt.Sprintf("UPDATE products SET %s WHERE id = ?", strings.Join(updates, ", "))
	_, err = s.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	return s.GetProductByID(productID)
}

func (s *Store) DeleteProductByID(productID int) error {
	result, err := s.db.Exec("DELETE FROM products WHERE id = ?", productID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("product with ID %d not found", productID)
	}

	return nil
}

func (s *Store) GetProducts() ([]*types.Product, error) {
	rows, err := s.db.Query("SELECT id, title, description, seller, rating, createdAt FROM products")
	if err != nil {
		return nil, err
	}

	var products []*types.Product
	for rows.Next() {
		product, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
