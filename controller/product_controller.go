package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/khaizbt/superindo-test/helper"
	"github.com/khaizbt/superindo-test/service"
	"net/http"
)

type product_controller struct {
	product_service service.ProductService
	user_service    service.UserService
}

func BankController(bankService service.ProductService, userService service.UserService) *product_controller {
	return &product_controller{bankService, userService}
}

func (h *product_controller) ListProducts(c *gin.Context) {
	products, err := h.product_service.ListProduct()

	if err != nil {
		responsError := helper.APIResponse("Product Not Found", http.StatusUnprocessableEntity, "fail", err)
		c.JSON(http.StatusUnprocessableEntity, responsError)
		return
	}
	response := helper.APIResponse("Success", http.StatusOK, "success", products)
	c.JSON(http.StatusOK, response)

}
