package service

import (
	"GoAlbums/internal/db"
)

func UpdateAccount(id string, updatedAccount User) (User, error) {
	var user User
	result := db.DB.DB.First(&user, "id = ?", id)
	if result.Error != nil {
		return User{}, result.Error
	}

	user.FullName = updatedAccount.FullName
	user.Email = updatedAccount.Email

	saveResult := db.DB.DB.Save(&user)
	if saveResult.Error != nil {
		return User{}, saveResult.Error
	}
	return user, nil
}
