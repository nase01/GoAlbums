package validator

import (
	"GoAlbums/internal/dto"
	"GoAlbums/utils/validator"
	"errors"
)

func ValidateUser(user dto.SignUpRequest) error {

	if !validator.ValidName(user.FullName) {
		return errors.New("invalid name format")
	}

	if !validator.ValidEmail(user.Email) {
		return errors.New("invalid email format")
	}

	if !validator.PasswordMatched(user.Password, user.ConfirmPassword) {
		return errors.New("password do not match")
	}

	if !validator.StrongPassword(user.Password) {
		return errors.New("password must be at least 8 characters long and contain at least one uppercase letter, one lowercase letter, one digit, and one special character")
	}

	return nil
}
