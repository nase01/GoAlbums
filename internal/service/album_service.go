package service

import (
	"GoAlbums/internal/db"
	"GoAlbums/internal/models"

	"gorm.io/gorm"
)

type Album models.Album

func GetAlbums(currentPage, perPage int, sort, from, to string) ([]Album, error) {
	var albums []Album

	offset := (currentPage - 1) * perPage
	if sort != "asc" && sort != "desc" {
		sort = "desc"
	}

	result := db.DB.DB.
		Where("created_at BETWEEN ? AND ?", from, to).
		Order("created_at " + sort).Limit(perPage).Offset(offset).Find(&albums)
	return albums, result.Error
}

func GetAlbumById(id string) (Album, error) {
	var album Album
	result := db.DB.DB.First(&album, "id = ?", id)
	return album, result.Error
}

func CreateAlbum(album Album) (Album, error) {
	result := db.DB.DB.Create(&album)
	return album, result.Error
}

func UpdateAlbum(id string, updatedAlbum Album) (Album, error) {
	var album Album
	result := db.DB.DB.First(&album, "id = ?", id)
	if result.Error != nil {
		return Album{}, result.Error
	}

	album.Title = updatedAlbum.Title
	album.Artist = updatedAlbum.Artist
	album.Price = updatedAlbum.Price

	saveResult := db.DB.DB.Save(&album)
	if saveResult.Error != nil {
		return Album{}, saveResult.Error
	}
	return album, nil
}

func DeleteAlbums(ids []string) error {
	result := db.DB.DB.Delete(&Album{}, "id IN (?)", ids)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
