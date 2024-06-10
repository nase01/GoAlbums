package handlers

import (
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !validator.ValidName(request.FullName) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid name format"})
		return
	}

	if !validator.ValidEmail(request.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	if !validator.PasswordMatched(request.Password, request.ConfirmPassword) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
		return
	}

	if !validator.StrongPassword(request.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be at least 6 characters long and contain at least one uppercase letter, one lowercase letter, one digit, and one special character"})
		return
	}

	hashedPassword, err := helpers.HashPassword(request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
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
