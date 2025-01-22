package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/khaizbt/superindo-test/config"
	"github.com/khaizbt/superindo-test/entity"
	"github.com/khaizbt/superindo-test/helper"
	"github.com/khaizbt/superindo-test/model"
	"github.com/khaizbt/superindo-test/service"
	"net/http"
)

type (
	userController struct {
		userService service.UserService
		authService config.AuthService
	}
)

func NewUserController(userService service.UserService, authService config.AuthService) *userController {
	return &userController{userService, authService}
}

type UserFormatter struct {
	UserID   string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func FormatUser(user model.User, token string) UserFormatter {
	formatter := UserFormatter{
		UserID:   user.ID,
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}

	return formatter
}

func (h *userController) CreateUser(c *gin.Context) {
	var input entity.DataUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		responseError := helper.APIResponse("Create Account Failed #U001", http.StatusUnprocessableEntity, "fail", err.Error())
		c.JSON(http.StatusUnprocessableEntity, responseError)
		return
	}

	errValidation := validation.ValidateStruct(&input,
		validation.Field(&input.Username, validation.Required, is.Alphanumeric),
		validation.Field(&input.Email, validation.Required, is.Email),
		validation.Field(&input.Password, validation.Required, validation.NilOrNotEmpty),
	)

	if errValidation != nil {
		responsError := helper.APIResponse("Create Account Failed #U001-1", http.StatusUnprocessableEntity, "fail", errValidation)
		c.JSON(http.StatusUnprocessableEntity, responsError)
		return
	}

	createUser, err := h.userService.CreateUser(input)

	if err != nil {
		responseError := helper.APIResponse("Create Account Failed #U002", http.StatusBadRequest, "fail", err.Error())
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", createUser)
	c.JSON(http.StatusOK, response)
}

func (h *userController) Login(c *gin.Context) {
	var input entity.LoginInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		responseError := helper.APIResponse("Login Failed #LOG001", http.StatusUnprocessableEntity, "fail", err.Error())
		c.JSON(http.StatusUnprocessableEntity, responseError)
		return
	}

	errValidation := validation.ValidateStruct(&input,
		validation.Field(&input.Email, validation.Required, is.Email),
		validation.Field(&input.Password, validation.Required, validation.NilOrNotEmpty),
	)

	if errValidation != nil {
		responsError := helper.APIResponse("Login Failed #LOG002-1", http.StatusUnprocessableEntity, "fail", errValidation)
		c.JSON(http.StatusUnprocessableEntity, responsError)
		return
	}

	loginUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		responseError := helper.APIResponse("Login Failed #LOG002", http.StatusUnprocessableEntity, "fail", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, responseError)
		return
	}

	token, err := h.authService.GenerateTokenUser(loginUser.ID)

	if err != nil {
		responsError := helper.APIResponse("Login Failed", http.StatusBadGateway, "fail", "Unable to generate token")
		c.JSON(http.StatusBadGateway, responsError)
		return
	}

	//For Set Header if login/logout

	response := helper.APIResponse("Login Success", http.StatusOK, "success", FormatUser(loginUser, token))

	c.JSON(http.StatusOK, response)
}

func (h *userController) LogOut(c *gin.Context) {

	session := sessions.Default(c) //FIXME sessions masih error(Tidak terbaca, padahal sudah diimport)
	session.Clear()
	session.Set("token", "")
	session.Options(sessions.Options{
		MaxAge: -1,
	})
	session.Save()

	response := helper.APIResponse("Logout Success", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)

}
