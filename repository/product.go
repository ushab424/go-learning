package repository

import (
	"database/sql"
	"errors"
)

type ProductRepository struct {
	db *sql.DB
}

// структура продукта
type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (p *ProductRepository) GetByID(id int) (*Product, error) {
	var prod Product
	err := p.db.QueryRow("SELECT id, name, price FROM products WHERE id = $1", id).Scan(&prod.ID, &prod.Name, &prod.Price)
	if err != nil {
		return nil, errors.New("product not found")
	}
	return &prod, nil
}

func (p *ProductRepository) Create(name string, price int) (int, error) {
	var ProductID int
	err := p.db.QueryRow("INSERT INTO products (name, price) VALUES ($1, $2) RETERNING id", name, price).Scan(&ProductID)
	if err != nil {
		return 0, errors.New("invalid request")
	}
	return ProductID, nil
}

func (p *ProductRepository) Delete(id int) error {
	_, err := p.db.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return errors.New("product not found")
	}
	return nil
}
