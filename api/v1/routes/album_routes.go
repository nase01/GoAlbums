package routes

import (
	"GoAlbums/api/v1/handlers"
	middleware "GoAlbums/internal"

	"github.com/gin-gonic/gin"
)

func SetupAlbumRoutes(router *gin.Engine) {
	api := router.Group("/api/v1", middleware.AuthRequired())
	{
		api.GET("/albums", handlers.GetAlbums)
		api.GET("/albums/:id", handlers.GetAlbumByID)
		api.POST("/albums", handlers.CreateAlbum)
		api.PUT("/albums/:id", handlers.UpdateAlbum)
		api.DELETE("/albums", handlers.DeleteAlbums)
	}
}
