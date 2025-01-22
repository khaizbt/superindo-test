package route

import (
	"github.com/gin-gonic/gin"
	"github.com/khaizbt/superindo-test/controller"
	"github.com/khaizbt/superindo-test/middleware"
	"github.com/khaizbt/superindo-test/service"
)

func RouteBank(route *gin.Engine, service service.ProductService, user_service service.UserService) {
	productController := controller.BankController(service, user_service)
	api := route.Group("/api/v1/product")

	api.GET("/list", productController.ListProducts)
	api.POST("/create", middleware.AuthMiddlewareUser(authService, user_service), productController.CreeateProduct)

}
