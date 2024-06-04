package routes

import (
	"GoAlbums/api/v1/handlers"

	"github.com/gin-gonic/gin"
)

func SetupAlbumRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.GET("/albums", handlers.GetAlbums)
		api.GET("/albums/:id", handlers.GetAlbumByID)
		api.POST("/albums", handlers.CreateAlbum)
		api.PUT("/albums/:id", handlers.UpdateAlbum)
		api.DELETE("/albums/:id", handlers.DeleteAlbum)
	}
}
