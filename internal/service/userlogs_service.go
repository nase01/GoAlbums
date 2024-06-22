package service

import (
	"GoAlbums/internal/db"
	"GoAlbums/internal/models"
)

type UserLogs models.UserLogs

func GetUserLogs() ([]UserLogs, error) {
	var userLogs []UserLogs

	result := db.DB.DB.Find(&userLogs)
	return userLogs, result.Error
}

func CreateUserLogs(userLogs UserLogs) (UserLogs, error) {
	result := db.DB.DB.Create(&userLogs)
	return userLogs, result.Error
}
