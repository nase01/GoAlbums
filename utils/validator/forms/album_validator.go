package validator

import (
	"GoAlbums/internal/service"
	"errors"
)

func ValidateAlbum(album service.Album) error {
	if album.Title == "" {
		return errors.New("title is empty")
	}

	if album.Artist == "" {
		return errors.New("artist is empty")
	}

	return nil
}
