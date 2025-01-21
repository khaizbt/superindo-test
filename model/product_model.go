package model

import "time"

type (
	Product struct {
		ID          string `gorm:"primary_key"`
		Name        string `gorm:"type:varchar(255);NOT NULL"`
		Stock       int    `gorm:"type:int(20);NOT NULL"`
		SKU         string `gorm:"type:varchar(255);NULL"`
		Description string `gorm:"type:text;NOT NULL"`
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
)
