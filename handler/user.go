package handler

import (
	"bisadonasi/helper"
	"bisadonasi/user"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func BuildErrorResponse(message string, err string) interface{} {
	splittedError := strings.Split(err, "\n")
	// res := Response{
	// 	Status:  false,
	// 	Message: message,
	// 	Errors:  splittedError,
	// 	Data:    data,
	// }
	return splittedError
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

	newUser, err := h.userService.RegisterUser(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	formatter := user.FormatUser(newUser, "token,,,,,")

	response := helper.ApiResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
