package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/khaizbt/superindo-test/entity"
	"github.com/khaizbt/superindo-test/helper"
	"github.com/khaizbt/superindo-test/service"
	"github.com/spf13/cast"
	"net/http"
	"regexp"
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
		fmt.Println(err)
		responsError := helper.APIResponse("Create product failed", http.StatusUnprocessableEntity, "fail", err)
		c.JSON(http.StatusUnprocessableEntity, responsError)
		return
	}

	errValidation := validation.ValidateStruct(&payload,
		validation.Field(&payload.Name, validation.Required),
		validation.Field(&payload.Price, validation.Required, validation.Min(1), validation.Max(100000000)),
		validation.Field(&payload.Stock, validation.Required, validation.Min(1), validation.Max(100000000)),
		validation.Field(&payload.Description, validation.Length(6, 255)),
		validation.Field(&payload.Category, validation.Required),
		validation.Field(&payload.SKU, validation.Match(regexp.MustCompile("^[a-zA-Z0-9_-]+$")).Error("must contain English letters digits, dash and underscore only"), validation.Length(3, 50)),
	)

	if errValidation != nil {

		responsError := helper.APIResponse("Create product failed", http.StatusUnprocessableEntity, "fail", errValidation)
		c.JSON(http.StatusUnprocessableEntity, responsError)
		return
	}

	errCategoryValidation := helper.GetCategory(payload.Category)

	if errCategoryValidation != nil {

		responsError := helper.APIResponse("Create product failed", http.StatusUnprocessableEntity, "fail", map[string]string{"error": fmt.Sprintf("%+v", errCategoryValidation)})
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

func (h *product_controller) GetCategory(c *gin.Context) {
	data, err := h.product_service.GetCategory()

	if err != nil {
		responsError := helper.APIResponse("get category failed", http.StatusBadRequest, "failed", err)
		c.JSON(http.StatusUnprocessableEntity, responsError)
		return
	}

	response := helper.APIResponse("success get category", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
