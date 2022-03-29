package handler

import (
	"bisadonasi/helper"
	"bisadonasi/user"
	"fmt"
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
	var req user.InputRegister

	err := c.ShouldBind(&req)

	if err != nil {
		error := helper.FormatValidationError(err)
		resError := gin.H{"errors": error}
		response := helper.ApiResponse("Register Failed", http.StatusUnprocessableEntity, "error", resError)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	checkEmail, err := h.userService.IsDuplicateEmail(req.Email)

	fmt.Println("cek email", checkEmail)

	if checkEmail == true {
		var resErrors []string
		// error := helper.FormatValidationError()
		resErrors = append(resErrors, "Email is Duplicated")
		resError := gin.H{"errors": resErrors}
		response := helper.ApiResponse("Register is failed", http.StatusBadRequest, "error", resError)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.RegisterUser(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	formatter := user.FormatUser(newUser, "token,,,,,")

	response := helper.ApiResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) LoginUser(c *gin.Context) {

	var req user.InputLogin

	err := c.ShouldBind(&req)

	if err != nil {
		error := helper.FormatValidationError(err)
		resError := gin.H{"errors": error}
		response := helper.ApiResponse("Login Failed", http.StatusUnprocessableEntity, "error", resError)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.LoginUser(req)

	if err != nil {
		resError := gin.H{"errors": err.Error()}

		response := helper.ApiResponse("Login Failed", http.StatusUnprocessableEntity, "Invalid Email/Password", resError)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, "token,,,,,")

	response := helper.ApiResponse("Success Login", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")

	if err != nil {
		data := gin.H{"isUploaded": false}
		response := helper.ApiResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	userId := 1

	path := fmt.Sprintf("images/%d-%s", userId, file.Filename)
	err = c.SaveUploadedFile(file, path)

	if err != nil {
		data := gin.H{"isUploaded": false}
		response := helper.ApiResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userId, path)

	if err != nil {
		data := gin.H{"isUploaded": false}
		response := helper.ApiResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"isUploaded": true}
	response := helper.ApiResponse("Success to upload avatar image", http.StatusOK, "success", data)
	c.JSON(http.StatusBadRequest, response)

}
