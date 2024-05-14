package util

import (
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

const (
	minPasswordLength  = 8
	requireLowercase   = true
	requireUppercase   = true
	requireDigit       = true
	requireSpecialChar = true
	specialChars       = "@$!%*?&"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

func CheckPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func IsStrongPassword(password string) bool {
	var hasLower, hasUpper, hasDigit, hasSpecial bool

	if len(password) < minPasswordLength {
		return false
	}

	for _, char := range password {
		switch {
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsDigit(char):
			hasDigit = true
		default:
			if strings.ContainsRune(specialChars, char) {
				hasSpecial = true
			}
		}

		if hasLower && hasUpper && hasDigit && hasSpecial {
			return true
		}
	}

	return false
}
