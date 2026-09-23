package service

import (
	"day20tier2/repository"
	"errors"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetProduct(id int) (*repository.Product, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) AddProduct(name string, price int) (int, error) {
	if name == "" {
		return 0, errors.New("name cannot be emoty")
	}
	if price <= 0 {
		return 0, errors.New("price must be greater than 0")
	}

	return s.repo.Create(name, price)
}

func (s *ProductService) RemoveProduct(id int) error {
	return s.repo.Delete(id)
}
