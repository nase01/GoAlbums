package service

import (
	"GoAlbums/internal/db"
	"GoAlbums/internal/models"
	"GoAlbums/utils/helpers"

	"gorm.io/gorm"
)

type UserLogs models.UserLogs

func GetUserLogs(filters helpers.QueryFilters) ([]UserLogs, error) {
	var userLogs []UserLogs
	var result *gorm.DB

	offset := (filters.Pagination.CurrentPage - 1) * filters.Pagination.PerPage
	if filters.Sort != "asc" && filters.Sort != "desc" {
		filters.Sort = "desc"
	}

	query := db.DB.DB.
		Where("created_at BETWEEN ? AND ?", filters.From, filters.To).
		Order("created_at " + filters.Sort).Limit(filters.Pagination.PerPage).Offset(offset)

	result = query.Find(&userLogs)
	return userLogs, result.Error
}

func CreateUserLogs(userLogs UserLogs) (UserLogs, error) {
	result := db.DB.DB.Create(&userLogs)
	return userLogs, result.Error
}
