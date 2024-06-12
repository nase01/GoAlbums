package validator

import (
	"regexp"
	"strings"
)

func ValidEmail(email string) bool {
	emailRegex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func PasswordMatched(password, confirmPassword string) bool {
	return password == confirmPassword
}

func ValidName(name string) bool {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		return false
	}
	for _, char := range name {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || char == ' ') {
			return false
		}
	}
	return true
}

func StrongPassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	hasUpper := false
	hasLower := false
	hasDigit := false
	hasSpecial := false
	specialChars := "`-=[]\\;',./~!@#$%^&*()_+{}|:\"<>?"

	for _, char := range password {
		if char >= 'A' && char <= 'Z' {
			hasUpper = true
		} else if char >= 'a' && char <= 'z' {
			hasLower = true
		} else if char >= '0' && char <= '9' {
			hasDigit = true
		} else if strings.ContainsRune(specialChars, char) {
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasDigit && hasSpecial
}
