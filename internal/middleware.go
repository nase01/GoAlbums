package middleware

import (
	"GoAlbums/routes/api/v1/handlers"
	"GoAlbums/utils/helpers"
	"errors"
	"log"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("MySecret123")

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			errorResponse, statusCode := helpers.CustomError(errors.New("unauthorized access"))
			c.JSON(statusCode, errorResponse)
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		claims := &handlers.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			errorResponse, statusCode := helpers.CustomError(errors.New("unauthorized access"))
			c.JSON(statusCode, errorResponse)
			c.Abort()
			return
		}

		log.Printf("Middleware Checkpoint")

		c.Set("userID", claims.ID)
		c.Set("userEmail", claims.Email)
		c.Next()
	}
}
