package handlers

import (
	"errors"
	"net/http"

	"GoAlbums/internal/service"
	"GoAlbums/utils/helpers"
	"GoAlbums/utils/validator"

	"github.com/gin-gonic/gin"
)

func UpdateAccount(c *gin.Context) {
	userID, exists := c.Get("userID") // This is the userID extracted from JWT
	if !exists {
		errorResponse, statusCode := helpers.CustomError(errors.New("userID not found"))
		c.JSON(statusCode, errorResponse)
		return
	}

	var updatedAccount service.User
	if err := c.BindJSON(&updatedAccount); err != nil {
		errorResponse, statusCode := helpers.CustomError(err)
		c.JSON(statusCode, errorResponse)
		return
	}

	if !validator.ValidName(updatedAccount.FullName) {
		errorResponse, statusCode := helpers.CustomError(errors.New("invalid name format"))
		c.JSON(statusCode, errorResponse)
		return
	}

	if !validator.ValidEmail(updatedAccount.Email) {
		errorResponse, statusCode := helpers.CustomError(errors.New("invalid email format"))
		c.JSON(statusCode, errorResponse)
		return
	}

	album, err := service.UpdateAccount(userID.(string), updatedAccount)
	if err != nil {
		errorResponse, statusCode := helpers.CustomError(err)
		c.JSON(statusCode, errorResponse)
		return
	}
	c.JSON(http.StatusOK, album)
}
