package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupPublicRoutes(router *gin.Engine) {
	api := router.Group("/")
	{
		api.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "Welcome to GoAlbums")
		})
	}
}
