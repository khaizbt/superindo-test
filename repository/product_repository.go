package repository

import (
	"github.com/khaizbt/superindo-test/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProductByID(productID string) (*model.Product, error)

	GetProducts() ([]model.Product, error)
}

func NewBankRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetProductByID(productID string) (*model.Product, error) {
	var product model.Product
	err := r.db.Where("id = ?", productID).Find(&product).Error

	if err != nil {
		return nil, err
	}

	return &product, nil
}
func (r *repository) GetProducts() ([]model.Product, error) {
	var results []model.Product

	err := r.db.Find(&results).Error

	return results, err
}
