package handlers

import (
	"net/http"

	"GoAlbums/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	id := c.Param("id")
	var updatedAlbum service.Album
	if err := c.BindJSON(&updatedAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	album, err := service.UpdateAlbum(id, updatedAlbum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, album)
}

func DeleteAlbums(c *gin.Context) {
	var ids []string
	if err := c.BindJSON(&ids); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := service.DeleteAlbums(ids)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record(s) Not Found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Album(s) deleted"})
}
