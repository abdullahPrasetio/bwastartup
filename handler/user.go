package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// tangkap input dari user
	// map input dari user ke struct RegisterUserInput
	// struct di atas kita passing sebagai parameter service

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {

		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponseError("register account failed", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponseError("register account failed", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(newUser, "ksdfjnsdklfnsld;sdfmls;dmf")

	response := helper.APIResponseSuccess("Account has been registered", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	// user memasukan input (email & password)
	// ?input ditangkap handler
	// mapping dari input ke input struct
	// input struct passing service
	// di service mencari dengan vabtuan repository user dengan email x
	// Mencocokan password
	var input user.LoginInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponseError("Authentication Failed", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedInUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponseError("Authentication Failed", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(loggedInUser, "askdnaskldmaslkdm")
	response := helper.APIResponseSuccess("Authentication success", formatter)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	// ada input email dari user
	// input email di mapping ke struct input
	// struct di input di passing ke service
	// service akan manggil repository - email sudah ada atau belum
	// repository -db
	var input user.CheckEmailInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponseError("Email checking failed", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	IsEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}

		response := helper.APIResponseError("Email checking failed", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	data := gin.H{
		"is_available": IsEmailAvailable,
	}
	metaMessage := "Email has been registered"
	if IsEmailAvailable {
		metaMessage = "Email is availables"
	}
	response := helper.APIResponseSuccess(metaMessage, data)
	c.JSON(http.StatusOK, response)
}
