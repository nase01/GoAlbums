package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {

			//Need to have authorization token validation here

			c.String(http.StatusUnauthorized, "Unauthorized Access")
			c.Abort()
			return
		}

		log.Printf("Middleware Checkpoint")

		// Add some rediirection logic here
		// c.Redirect(http.StatusFound, "/dashboard")

		c.Next()
	}
}
