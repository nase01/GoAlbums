package routes

import (
	"GoAlbums/routes/api/v1/handlers"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.POST("/sign-up", handlers.SignUp)
	}
}
