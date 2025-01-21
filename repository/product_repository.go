package repository

import (
	"fmt"
	"github.com/khaizbt/superindo-test/entity"
	"github.com/khaizbt/superindo-test/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProductByID(productID string) (*model.Product, error)

	GetProducts(where entity.ProductQuery) ([]model.Product, error)
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
func (r *repository) GetProducts(where entity.ProductQuery) ([]model.Product, error) {
	var results []model.Product

	q := r.db.Select("products.id", "products.name", "GROUP_CONCAT(categories.name) AS product_category").
		Joins("JOIN product_categories ON products.id = product_categories.product_id").
		Joins("JOIN  categories  ON product_categories.category_id = categories.id").Limit(10).Offset(0)

	if where.ID != "" {
		q.Where("product_categories.product_id = ?", where.ID)
	}

	if where.Name != "" {
		q.Where("products.name LIKE ?", fmt.Sprintf("%%%s%%", where.Name))
	}

	if len(where.Category) > 0 {

		q.Where("categories.id IN ?", where.Category)

	}

	err := q.Group("products.id").Find(&results).Error

	return results, err
}
