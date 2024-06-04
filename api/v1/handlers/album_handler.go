package handlers

import (
	"net/http"

	"GoAlbums/internal/service"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	albums, err := service.GetAlbums()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, albums)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	album, err := service.GetAlbumById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record Not Found"})
		return
	}
	c.JSON(http.StatusOK, album)
}

func CreateAlbum(c *gin.Context) {
	var newAlbum service.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	album, err := service.CreateAlbum(newAlbum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, album)
}

func UpdateAlbum(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "Not Implemented"})
}

func DeleteAlbum(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "Not Implemented"})
}
