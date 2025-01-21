package main

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/khaizbt/superindo-test/middleware"
	"github.com/khaizbt/superindo-test/repository"
	"github.com/khaizbt/superindo-test/route"
	"github.com/khaizbt/superindo-test/service"
)

func main() {
	repo := repository.NewRepository()
	userService := service.NewUserService(repo)
	bankService := service.NewProductService(repo, sync.Mutex{})

	router := gin.Default()
	router.Use(middleware.SecureMiddleware())
	route.RouteUser(router, userService)
	route.RouteBank(router, bankService, userService)

	fmt.Println("starting server on port :8000")

	_ = router.Run(":8000")

}
