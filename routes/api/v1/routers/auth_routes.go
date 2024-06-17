package routes

import (
	middleware "GoAlbums/internal"
	"GoAlbums/routes/api/v1/handlers"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.POST("/auth/signin", handlers.SignIn)
		api.POST("/auth/signup", handlers.SignUp)

		authGroup := api.Group("")
		authGroup.Use(middleware.AuthRequired())
		{
			authGroup.GET("/auth/user", handlers.GetCurrentUser)
			// add more routes here
		}

	}
}
