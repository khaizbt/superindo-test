package route

import (
	"github.com/gin-gonic/gin"
	"github.com/khaizbt/superindo-test/config"
	"github.com/khaizbt/superindo-test/controller"
	"github.com/khaizbt/superindo-test/middleware"
	"github.com/khaizbt/superindo-test/service"
	"time"
)

var authService = config.NewServiceAuth()

func RouteUser(route *gin.Engine, service service.UserService) {
	userController := controller.NewUserController(service, authService)
	api := route.Group("/api/v1/user")
	api.POST("/register", userController.CreateUser)
	api.POST("/login", userController.Login)
	api.POST("/logout", middleware.AuthMiddlewareUser(authService, service), userController.LogOut)
	api.POST("/get-user", func(context *gin.Context) {

		context.JSON(200, gin.H{
			"message": "pong",
		})
		time.Sleep(1 * time.Second)
		return
	})

	api.POST("/all-user", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "success",
		})
	})
}
