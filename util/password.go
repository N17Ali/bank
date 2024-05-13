package util

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
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

const (
	minPasswordLength  = 8
	requireLowercase   = true
	requireUppercase   = true
	requireDigit       = true
	requireSpecialChar = true
	specialChars       = "@$!%*?&"
)

func IsStrongPassword(password string) bool {
	var (
		hasLower, hasUpper, hasDigit, hasSpecial bool
	)

	if len(password) < minPasswordLength {
		return false
	}

	for _, c := range password {
		switch {
		case c >= 'a' && c <= 'z':
			hasLower = true
		case c >= 'A' && c <= 'Z':
			hasUpper = true
		case c >= '0' && c <= '9':
			hasDigit = true
		default:
			hasSpecial = strings.ContainsRune(specialChars, c)
		}

		if hasLower && hasUpper && hasDigit && hasSpecial {
			return true
		}
	}

	if requireLowercase && !hasLower {
		return false
	}
	if requireUppercase && !hasUpper {
		return false
	}
	if requireDigit && !hasDigit {
		return false
	}
	if requireSpecialChar && !hasSpecial {
		return false
	}

	return true
}
