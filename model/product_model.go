package model

import "time"

type (
	Product struct {
		ID              string    `gorm:"primary_key" json:"id"`
		Name            string    `gorm:"type:varchar(255);NOT NULL" json:"name"`
		Stock           int       `gorm:"type:int(20);NOT NULL" json:"stock"`
		SKU             string    `gorm:"type:varchar(255);NULL" json:"sku"`
		Description     *string   `gorm:"type:text;NOT NULL" json:"description"`
		ProductCategory string    `json:"product_category"`
		CreatedAt       time.Time `json:"created_at"`
		UpdatedAt       time.Time `json:"updated_at"`
	}

	ProductCategory struct {
		ID         string `gorm:"primary_key"`
		ProductID  string `gorm:"primary_key"`
		CategoryID string `gorm:"primary_key"`
		CreatedAt  time.Time
		UpdatedAt  time.Time
	}

	Category struct {
		ID          string `gorm:"primary_key"`
		Name        string `gorm:"type:varchar(255);NOT NULL"`
		Description string `gorm:"type:text;NOT NULL"`
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
)
