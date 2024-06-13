package service

import (
	"GoAlbums/internal/db"
	"GoAlbums/internal/models"

	"gorm.io/gorm"
)

type User models.User

func SignIn(db *gorm.DB, email, password string) (*models.User, error) {
	user, err := models.FindUserByEmail(db, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func SignUp(user User) (User, error) {
	result := db.DB.DB.Create(&user)
	return user, result.Error
}
