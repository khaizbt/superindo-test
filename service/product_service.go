package service

import (
	"github.com/google/uuid"
	"github.com/khaizbt/superindo-test/entity"
	"github.com/spf13/cast"
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
		CreateProduct(query entity.ProductCreateInput) error
		ListProduct(query entity.ProductQuery) ([]model.ProductList, error)
		GetCategory() ([]model.Category, error)
	}
)

func NewProductService(product_repository repository.ProductRepository, mu sync.Mutex) *product_service {
	return &product_service{product_repository, mu}
}

func (s *product_service) ListProduct(query entity.ProductQuery) ([]model.ProductList, error) {
	if len(query.Category) == 1 {
		for _, category := range query.Category {
			if category == "" {
				query.Category = nil
				break
			}
		}
	}

	if query.PerPage == 0 {
		query.PerPage = 10
	}

	query.Page = (query.Page - 1) * query.PerPage
	product, err := s.repository.GetProducts(query)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *product_service) CreateProduct(query entity.ProductCreateInput) error {
	productId := uuid.NewString()

	errCreateProduct := s.repository.Create(model.Product{
		ID:          productId,
		Name:        query.Name,
		Stock:       cast.ToInt(query.Stock),
		SKU:         query.SKU,
		Description: &query.Description,
		Price:       cast.ToInt(query.Price),
	})

	if errCreateProduct != nil {
		return errCreateProduct
	}

	//get product category
	categories, err := s.repository.GetProductCategory(query.Category)

	if err != nil {
		return err
	}

	for _, category := range categories {
		errCreateProductCategory := s.repository.CreateProductCategory(model.ProductCategory{
			ID:         uuid.NewString(),
			ProductID:  productId,
			CategoryID: category.ID,
		})

		if errCreateProductCategory != nil {
			return errCreateProductCategory
		}
	}

	return nil

}

func (s *product_service) GetCategory() ([]model.Category, error) {
	result, err := s.repository.GetCategory()

	if err != nil {
		return nil, err
	}

	return result, nil
}
