package middleware

import (
	"GoAlbums/config"
	"GoAlbums/routes/api/v1/handlers"
	"GoAlbums/utils/helpers"
	"errors"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = config.GetJWTKey()

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
		c.Set("userRole", claims.Role)
		c.Next()
	}
}
