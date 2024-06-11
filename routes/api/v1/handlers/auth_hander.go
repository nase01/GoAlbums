package handlers

import (
	"errors"
	"net/http"

	"GoAlbums/internal/dto"
	"GoAlbums/internal/models"
	"GoAlbums/internal/service"
	"GoAlbums/utils/helpers"
	"GoAlbums/utils/validator"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var request dto.SignUpRequest

	if err := c.BindJSON(&request); err != nil {
		errorResponse, statusCode := helpers.CustomError(err)
		c.JSON(statusCode, errorResponse)
		return
	}

	if !validator.ValidName(request.FullName) {
		errorResponse, statusCode := helpers.CustomError(errors.New("invalid name format"))
		c.JSON(statusCode, errorResponse)
		return
	}

	if !validator.ValidEmail(request.Email) {
		errorResponse, statusCode := helpers.CustomError(errors.New("invalid email format"))
		c.JSON(statusCode, errorResponse)
		return
	}

	if !validator.PasswordMatched(request.Password, request.ConfirmPassword) {
		errorResponse, statusCode := helpers.CustomError(errors.New("password do not match"))
		c.JSON(statusCode, errorResponse)
		return
	}

	if !validator.StrongPassword(request.Password) {
		errorResponse, statusCode := helpers.CustomError(errors.New("password must be at least 6 characters long and contain at least one uppercase letter, one lowercase letter, one digit, and one special character"))
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
