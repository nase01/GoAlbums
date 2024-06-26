package service

import (
	"GoAlbums/internal/db"
	"GoAlbums/internal/models"
)

type UserLogs models.UserLogs

func GetUserLogs(currentPage, perPage int, sort string) ([]UserLogs, error) {
	var userLogs []UserLogs

	offset := (currentPage - 1) * perPage
	if sort != "asc" && sort != "desc" {
		sort = "desc"
	}

	result := db.DB.DB.Order("created_at " + sort).Limit(perPage).Offset(offset).Find(&userLogs)
	return userLogs, result.Error
}

func CreateUserLogs(userLogs UserLogs) (UserLogs, error) {
	result := db.DB.DB.Create(&userLogs)
	return userLogs, result.Error
}
