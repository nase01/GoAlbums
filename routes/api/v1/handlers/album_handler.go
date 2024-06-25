package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"GoAlbums/internal/service"
	"GoAlbums/utils/helpers"
	form "GoAlbums/utils/validator/forms"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {

	pagination := helpers.GetPaginationParams(c)

	albums, err := service.GetAlbums(pagination.CurrentPage, pagination.PerPage)
	if err != nil {
		errorResponse, statusCode := helpers.CustomError(err)
		c.JSON(statusCode, errorResponse)
		return
	}
	c.JSON(http.StatusOK, albums)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	album, err := service.GetAlbumById(id)
	if err != nil {
		errorResponse, statusCode := helpers.CustomError(err)
		c.JSON(statusCode, errorResponse)
		return
	}
	c.JSON(http.StatusOK, album)
}

func CreateAlbum(c *gin.Context) {
	var newAlbum service.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		errorResponse, statusCode := helpers.CustomError(err)
		c.JSON(statusCode, errorResponse)
		return
	}

	if err := form.ValidateAlbum(newAlbum); err != nil {
		errorResponse, statusCode := helpers.CustomError(err)
		c.JSON(statusCode, errorResponse)
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		errorResponse, statusCode := helpers.CustomError(errors.New("userID not found"))
		c.JSON(statusCode, errorResponse)
		return
	}

	newAlbum.UpdatedBy = userID.(string)

	album, err := service.CreateAlbum(newAlbum)
	if err != nil {
		errorResponse, statusCode := helpers.CustomError(err)
		c.JSON(statusCode, errorResponse)
		return
	}

	userLogs := service.UserLogs{
		UserId:   userID.(string),
		Activity: fmt.Sprintf("Created album -> %v", album),
	}

	if _, err := service.CreateUserLogs(userLogs); err != nil {
		log.Printf("Failed to log user activity: %v", err)
	}

	c.JSON(http.StatusCreated, album)
}

func UpdateAlbum(c *gin.Context) {

	id := c.Param("id")
	oldAlbum, _ := service.GetAlbumById(id)

	var updatedAlbum service.Album
	if err := c.BindJSON(&updatedAlbum); err != nil {
		errorResponse, statusCode := helpers.CustomError(err)
		c.JSON(statusCode, errorResponse)
		return
	}

	if err := form.ValidateAlbum(updatedAlbum); err != nil {
		errorResponse, statusCode := helpers.CustomError(err)
		c.JSON(statusCode, errorResponse)
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		errorResponse, statusCode := helpers.CustomError(errors.New("userID not found"))
		c.JSON(statusCode, errorResponse)
		return
	}

	updatedAlbum.UpdatedBy = userID.(string)

	album, err := service.UpdateAlbum(id, updatedAlbum)
	if err != nil {
		errorResponse, statusCode := helpers.CustomError(err)
		c.JSON(statusCode, errorResponse)
		return
	}

	userLogs := service.UserLogs{
		UserId:   userID.(string),
		Activity: fmt.Sprintf("Modified album -> %v > %v", oldAlbum, album),
	}

	if _, err := service.CreateUserLogs(userLogs); err != nil {
		log.Printf("Failed to log user activity: %v", err)
	}

	c.JSON(http.StatusOK, album)
}

func DeleteAlbums(c *gin.Context) {
	var ids []string
	if err := c.BindJSON(&ids); err != nil {
		errorResponse, statusCode := helpers.CustomError(err)
		c.JSON(statusCode, errorResponse)
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		errorResponse, statusCode := helpers.CustomError(errors.New("userID not found"))
		c.JSON(statusCode, errorResponse)
		return
	}

	err := service.DeleteAlbums(ids)
	if err != nil {
		errorResponse, statusCode := helpers.CustomError(err)
		c.JSON(statusCode, errorResponse)
		return
	}

	for _, id := range ids {
		userLog := service.UserLogs{
			UserId:   userID.(string),
			Activity: fmt.Sprintf("Deleted album -> %s", id),
		}

		if _, err := service.CreateUserLogs(userLog); err != nil {
			log.Printf("Failed to log user activity for album ID %s: %v", id, err)
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Album(s) deleted"})
}
