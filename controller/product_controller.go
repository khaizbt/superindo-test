package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/khaizbt/superindo-test/entity"
	"github.com/khaizbt/superindo-test/helper"
	"github.com/khaizbt/superindo-test/service"
	"github.com/spf13/cast"
	"net/http"
	"strings"
)

type product_controller struct {
	product_service service.ProductService
	user_service    service.UserService
}

func BankController(bankService service.ProductService, userService service.UserService) *product_controller {
	return &product_controller{bankService, userService}
}

func (h *product_controller) ListProducts(c *gin.Context) {
	var payload entity.ProductQuery
	payload.ID = c.Query("product_id")
	payload.Name = c.Query("product_name")
	productCategory := c.Query("product_category")
	payload.OrderBy = c.Query("order_by")
	payload.OrderMethod = c.Query("order_method")

	perPage := c.Query("per_page")
	page := c.Query("page")

	payload.PerPage = cast.ToInt(perPage)
	payload.Page = cast.ToInt(page)

	payload.Category = strings.Split(productCategory, ",")

	products, err := h.product_service.ListProduct(payload)

	if err != nil {
		responsError := helper.APIResponse("Product Not Found", http.StatusUnprocessableEntity, "fail", err)
		c.JSON(http.StatusUnprocessableEntity, responsError)
		return
	}
	response := helper.APIResponse("Success", http.StatusOK, "success", products)
	c.JSON(http.StatusOK, response)

}

func (h *product_controller) CreeateProduct(c *gin.Context) {
	var payload entity.ProductCreateInput

	err := c.ShouldBindJSON(&payload)

	if err != nil {
		responsError := helper.APIResponse("Create product failed", http.StatusUnprocessableEntity, "fail", err)
		c.JSON(http.StatusUnprocessableEntity, responsError)
		return
	}

	errCreate := h.product_service.CreateProduct(payload)

	if errCreate != nil {
		responsError := helper.APIResponse("Create product failed", http.StatusBadRequest, "fail", err)
		c.JSON(http.StatusUnprocessableEntity, responsError)
		return
	}

	response := helper.APIResponse("Product has been created", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)

}
