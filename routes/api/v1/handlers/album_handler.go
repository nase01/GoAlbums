package handlers

import (
	"net/http"

	"GoAlbums/internal/service"
	"GoAlbums/utils/helpers"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	albums, err := service.GetAlbums()
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

	album, err := service.CreateAlbum(newAlbum)
	if err != nil {
		errorResponse, statusCode := helpers.CustomError(err)
		c.JSON(statusCode, errorResponse)
		return
	}
	c.JSON(http.StatusCreated, album)
}

func UpdateAlbum(c *gin.Context) {
	id := c.Param("id")
	var updatedAlbum service.Album
	if err := c.BindJSON(&updatedAlbum); err != nil {
		errorResponse, statusCode := helpers.CustomError(err)
		c.JSON(statusCode, errorResponse)
		return
	}

	album, err := service.UpdateAlbum(id, updatedAlbum)
	if err != nil {
		errorResponse, statusCode := helpers.CustomError(err)
		c.JSON(statusCode, errorResponse)
		return
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

	err := service.DeleteAlbums(ids)
	if err != nil {
		errorResponse, statusCode := helpers.CustomError(err)
		c.JSON(statusCode, errorResponse)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Album(s) deleted"})
}
