package service

import (
	"github.com/khaizbt/superindo-test/entity"
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
		ListProduct(query entity.ProductQuery) ([]model.Product, error)
	}
)

func NewProductService(product_repository repository.ProductRepository, mu sync.Mutex) *product_service {
	return &product_service{product_repository, mu}
}

func (s *product_service) ListProduct(query entity.ProductQuery) ([]model.Product, error) {
	if len(query.Category) == 1 {
		for _, category := range query.Category {
			if category == "" {
				query.Category = nil
				break
			}
		}
	}
	product, err := s.repository.GetProducts(query)

	if err != nil {
		return nil, err
	}

	return product, nil
}
