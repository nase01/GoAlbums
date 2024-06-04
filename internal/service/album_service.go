package service

import (
	"GoAlbums/internal/db"
	"GoAlbums/internal/models"
)

type Album models.Album

func GetAlbums() ([]Album, error) {
	var albums []Album

	result := db.DB.DB.Find(&albums)
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
