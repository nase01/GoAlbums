package middleware

import (
	"GoAlbums/utils/helpers"
	"errors"
	"log"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {

			//Need to have authorization token validation here

			errorResponse, statusCode := helpers.CustomError(errors.New("unauthorized access"))
			c.JSON(statusCode, errorResponse)
			c.Abort()
			return
		}

		log.Printf("Middleware Checkpoint")

		// Add some rediirection logic here
		// c.Redirect(http.StatusFound, "/dashboard")

		c.Next()
	}
}
