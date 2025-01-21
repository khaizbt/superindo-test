package service

import (
	"sync"

	"github.com/khaizbt/superindo-test/model"
	"github.com/khaizbt/superindo-test/repository"
)

type (
	product_service struct {
		repository repository.ProductRepository
		mu         sync.Mutex
	}

	ProductService interface {
		ListProduct() ([]model.Product, error)
	}
)

func NewProductService(product_repository repository.ProductRepository, mu sync.Mutex) *product_service {
	return &product_service{product_repository, mu}
}

func (s *product_service) ListProduct() ([]model.Product, error) {
	product, err := s.repository.GetProducts()

	if err != nil {
		return nil, err
	}

	return product, nil
}
