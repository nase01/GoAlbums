package routes

import (
	middleware "GoAlbums/internal"
	"GoAlbums/routes/api/v1/handlers"

	"github.com/gin-gonic/gin"
)

func SetupAlbumRoutes(router *gin.Engine) {
	api := router.Group("/api/v1", middleware.AuthRequired())
	{
		api.GET("/albums", handlers.GetAlbums)
		api.GET("/albums/:id", handlers.GetAlbumByID)
		api.POST("/albums", middleware.RoleRequired("super"), handlers.CreateAlbum)
		api.PUT("/albums/:id", middleware.RoleRequired("super"), handlers.UpdateAlbum)
		api.DELETE("/albums", middleware.RoleRequired("super"), handlers.DeleteAlbums)
	}
}
