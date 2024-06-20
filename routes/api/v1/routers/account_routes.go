package routes

import (
	middleware "GoAlbums/internal"
	"GoAlbums/routes/api/v1/handlers"

	"github.com/gin-gonic/gin"
)

func SetupAccountRoutes(router *gin.Engine) {
	api := router.Group("/api/v1", middleware.AuthRequired())
	{
		api.PUT("/account", handlers.UpdateAccount)
	}
}
