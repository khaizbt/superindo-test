package repository

import (
	"fmt"
	"github.com/khaizbt/superindo-test/entity"
	"github.com/khaizbt/superindo-test/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProductByID(productID string) (*model.Product, error)

	GetProducts(where entity.ProductQuery) ([]model.ProductList, error)
	Create(product model.Product) error
	CreateProductCategory(productCategory model.ProductCategory) error
	GetProductCategory(category []string) ([]model.Category, error)
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
func (r *repository) GetProducts(where entity.ProductQuery) ([]model.ProductList, error) {
	var results []model.ProductList

	q := r.db.Table("products").Select("products.id", "products.name", "products.stock", "products.price", "products.sku", "products.created_at", "GROUP_CONCAT(categories.name) AS product_category").
		Joins("JOIN product_categories ON products.id = product_categories.product_id").
		Joins("JOIN  categories  ON product_categories.category_id = categories.id")

	if where.ID != "" {
		q.Where("product_categories.product_id = ?", where.ID)
	}

	if where.Name != "" {
		q.Where("products.name LIKE ?", fmt.Sprintf("%%%s%%", where.Name))
	}

	if len(where.Category) > 0 {

		q.Where("categories.id IN ?", where.Category)

	}

	q.Group("products.id")

	if where.OrderBy != "" {
		q.Order(fmt.Sprintf("%s %s", where.OrderBy, where.OrderMethod))
	} else {
		q.Order(fmt.Sprintf("%s %s", "name", "asc"))
	}

	err := q.Limit(where.PerPage).Offset(where.Page).Find(&results).Error

	return results, err
}

func (r *repository) Create(product model.Product) error {
	err := r.db.Create(&product).Error

	return err
}

func (r *repository) CreateProductCategory(productCategory model.ProductCategory) error {
	err := r.db.Create(&productCategory).Error

	return err
}

func (r *repository) GetProductCategory(category []string) ([]model.Category, error) {
	var results []model.Category

	err := r.db.Where("id IN ?", category).Find(&results).Error

	return results, err
}
