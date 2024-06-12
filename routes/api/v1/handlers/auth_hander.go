package handlers

import (
	"net/http"

	"GoAlbums/internal/dto"
	"GoAlbums/internal/models"
	"GoAlbums/internal/service"
	"GoAlbums/utils/helpers"
	form "GoAlbums/utils/validator/forms"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var request dto.SignUpRequest

	if err := c.BindJSON(&request); err != nil {
		errorResponse, statusCode := helpers.CustomError(err)
		c.JSON(statusCode, errorResponse)
		return
	}

	if err := form.ValidateUser(request); err != nil {
		errorResponse, statusCode := helpers.CustomError(err)
		c.JSON(statusCode, errorResponse)
		return
	}

	hashedPassword, err := helpers.HashPassword(request.Password)
	if err != nil {
		errorResponse, statusCode := helpers.CustomError(err)
		c.JSON(statusCode, errorResponse)
		return
	}
	request.Password = hashedPassword

	newUser := models.User{
		Email:    request.Email,
		Password: hashedPassword,
		FullName: request.FullName,
	}

	user, err := service.SignUp(service.User(newUser))
	if err != nil {
		errorResponse, statusCode := helpers.CustomError(err)
		c.JSON(statusCode, errorResponse)
		return
	}

	c.JSON(http.StatusCreated, user)
}
