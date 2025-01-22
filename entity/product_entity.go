package entity

import (
	"time"
)

type (
	AddToCartInput struct {
		ProductID string `json:"product_id" binding:"required"`
		Amount    int    `json:"amount" binding:"required"`
	}
	CheckoutInput struct {
		CartID string `json:"cart_id" binding:"required"`
	}

	OrderResponse struct {
		ID          string    `json:"id"`
		ProductName string    `json:"product_name"`
		Amount      string    `json:"amount"`
		CreatedAt   time.Time `json:"created_at"`
	}

	ProductQuery struct {
		Name        string   `form:"name"`
		ID          string   `form:"id"`
		Category    []string `form:"category"`
		OrderBy     string   `form:"order_by"`
		OrderMethod string   `form:"order_method"`
		PerPage     int      `form:"per_page"`
		Page        int      `form:"page"`
	}

	ProductCreateInput struct {
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Stock       int      `json:"stock"`
		Price       int      `json:"price"`
		Category    []string `json:"category"`
		SKU         string   `json:"sku"`
	}
)
