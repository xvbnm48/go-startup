package handler

import (
	"fmt"
	"go-startup/auth"
	"go-startup/helper"
	"go-startup/user"

	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}

}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// tangkap input dari user
	// map input dari struct RegisterUserINput
	// struct di atas kita parsing sebagai parameter

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("registered account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("registered account failed", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.APIResponse("registered account failed", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return

	}

	formatter := user.FormatUser(newUser, token)

	response := helper.APIResponse("account has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) Login(c *gin.Context) {
	// user memasukkann input email dan password
	// input di tangkap handler
	// mapping dari input ke struct
	// input struct parsing ke service
	// di service mencari dengan bantuan repository email yang di inputkan
	// mencocokkan password

	var input user.LoginInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	loggedInUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	token, err := h.authService.GenerateToken(loggedInUser.ID)
	if err != nil {
		response := helper.APIResponse("Loginn failed", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return

	}

	formatter := user.FormatUser(loggedInUser, token)
	response := helper.APIResponse("successfuly login", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	// input email dari user
	// input email di mapping ke struct input
	// struct di parsing ke service
	// dari service memanggil repository untuk menentukan email sudah ada atau belum
	// repo query ke database

	var input user.CheckEmailInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("email check failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}
	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "server error"}

		response := helper.APIResponse("email check failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	metaMessage := "email is registered"
	if isEmailAvailable {
		metaMessage = "email is available"
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}
	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
	return

}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	// untuk menerima inputan dari header berupa file
	file, err := c.FormFile("avatar")

	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse("failed upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//dari jwt tapi sekarang hard code
	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.ID

	//path := "images/" + file.Filename
	path := fmt.Sprintf("images/%d-%s", userId, file.Filename)

	err = c.SaveUploadedFile(file, path)

	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse("failed upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	_, err = h.userService.SaveAvatar(userId, path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse("failed upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{
		"is_uploaded": true,
	}
	response := helper.APIResponse("avatar image success uploaded", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}
func (h *userHandler) FetchUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	formatter := user.FormatUser(currentUser, "")
	response := helper.APIResponse("Successfully fetch user data", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}
