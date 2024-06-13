package models

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	CtmBasemModel
	Email    string `json:"email" gorm:"column:email;unique;index"`
	Password string `json:"-" gorm:"column:password;varchar(255);"` // Exclude Password from JSON Output
	FullName string `json:"fullname" gorm:"column:fullname;not null;varchar(255);"`
}

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
