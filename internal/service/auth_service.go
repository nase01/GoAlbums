package service

import (
	"GoAlbums/internal/db"
	"GoAlbums/internal/models"
	"errors"

	"gorm.io/gorm"
)

type User models.User

func FindUserByEmail(db *gorm.DB, email string) (*User, error) {
	var user User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func SignUp(user User) (User, error) {
	result := db.DB.DB.Create(&user)
	return user, result.Error
}
