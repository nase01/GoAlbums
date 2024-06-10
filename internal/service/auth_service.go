package service

import (
	"GoAlbums/internal/db"
	"GoAlbums/internal/models"
)

type User models.User

func SignUp(user User) (User, error) {
	result := db.DB.DB.Create(&user)
	return user, result.Error
}
