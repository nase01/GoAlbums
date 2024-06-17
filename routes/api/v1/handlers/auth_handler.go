package handlers

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"GoAlbums/config"
	"GoAlbums/internal/db"
	"GoAlbums/internal/dto"
	"GoAlbums/internal/models"
	"GoAlbums/internal/service"
	"GoAlbums/utils/helpers"
	form "GoAlbums/utils/validator/forms"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = config.GetJWTKey()

type Claims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
	IP string `json:"ip"`
}

func SignIn(c *gin.Context) {
	var request dto.SignInRequest
	if err := c.BindJSON(&request); err != nil {
		errorResponse, statusCode := helpers.CustomError(err)
		c.JSON(statusCode, errorResponse)
		return
	}

	user, err := service.FindUserByEmail(db.DB.DB, request.Email)
	if err != nil {
		errorResponse, statusCode := helpers.CustomError(err)
		c.JSON(statusCode, errorResponse)
		return
	}

	log.Printf("Role: %v", string(user.Role))
	log.Printf("IP: %v", helpers.GetUserIP(c.Request))

	if !helpers.CheckPasswordHash(request.Password, user.Password) {
		errorResponse, statusCode := helpers.CustomError(errors.New("invalid password"))
		c.JSON(statusCode, errorResponse)
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		ID:    user.Id,
		Email: user.Email,
		IP:    helpers.GetUserIP(c.Request),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		errorResponse, statusCode := helpers.CustomError(err)
		c.JSON(statusCode, errorResponse)
		return
	}

	response := map[string]interface{}{
		"data": map[string]interface{}{
			"success": true,
			"user": map[string]interface{}{
				"id":    user.Id,
				"email": user.Email,
				"token": tokenString,
			},
		},
	}
	c.JSON(http.StatusOK, response)
}

func SignUp(c *gin.Context) {
	var request dto.SignUpRequest

	if err := c.BindJSON(&request); err != nil {
		errorResponse, statusCode := helpers.CustomError(err)
		c.JSON(statusCode, errorResponse)
		return
	}

	if err := form.ValidateUser(request); err != nil {
		errorResponse, statusCode := helpers.CustomError(err)
		c.JSON(statusCode, errorResponse)
		return
	}

	hashedPassword, err := helpers.HashPassword(request.Password)
	if err != nil {
		errorResponse, statusCode := helpers.CustomError(err)
		c.JSON(statusCode, errorResponse)
		return
	}
	request.Password = hashedPassword

	newUser := models.User{
		Email:    request.Email,
		Password: hashedPassword,
		FullName: request.FullName,
	}

	user, err := service.SignUp(service.User(newUser))
	if err != nil {
		errorResponse, statusCode := helpers.CustomError(err)
		c.JSON(statusCode, errorResponse)
		return
	}

	c.JSON(http.StatusCreated, user)
}

func GetCurrentUser(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	tokenString := strings.Split(authHeader, " ")[1]

	parts := strings.Split(tokenString, ".")
	payload, _ := base64.RawStdEncoding.DecodeString(parts[1])

	claims := &Claims{}
	json.Unmarshal(payload, claims)

	userID := claims.ID

	user, err := service.FindUserByID(db.DB.DB, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
