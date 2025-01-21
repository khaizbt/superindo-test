package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/khaizbt/superindo-test/config"
	"github.com/khaizbt/superindo-test/helper"
	"github.com/khaizbt/superindo-test/service"
)

func AuthMiddlewareUser(authService config.AuthService, service service.UserService) gin.HandlerFunc {
	return func(context *gin.Context) {

		authHeader := context.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized #TKN001", http.StatusUnauthorized, "error", nil)
			context.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		var tokenString string
		arrayToken := strings.Split(authHeader, " ")

		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)

		if err != nil {
			response := helper.APIResponse("Unauthorized #TKN002", http.StatusUnauthorized, "error", nil)
			context.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized #TKN003", http.StatusUnauthorized, "error", nil)
			context.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := claim["user_id"].(string)

		//Check Redis, jika user id ada maka ambil dari redis, misal tidak ada maka ambil dari db lalu set ke redis selama 2 jam
		user, err := service.GetUserById(userID) //Jadikan Dari Redis

		if err != nil {
			response := helper.APIResponse("Unauthorized #TKN004", http.StatusUnauthorized, "error", nil)
			context.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		context.Set("loggedUser", user)
	}
}
