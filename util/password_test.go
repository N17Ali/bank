package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(6)
	hashedPassword1, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword1)

	err = CheckPassword(password, hashedPassword1)
	require.NoError(t, err)

	hashedPassword2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword2)
	require.NotEqual(t, hashedPassword1, hashedPassword2)

	wrongPassword := RandomString(6)
	err = CheckPassword(wrongPassword, hashedPassword1)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

}

func TestRandomPassword(t *testing.T) {
	password := RandomPassword()

	// Test that the generated password is at least the minimum length
	require.GreaterOrEqual(t, len(password), minPasswordLength)

	// Test that the generated password is valid
	require.True(t, IsStrongPassword(password))

	// Test that the generated password is different each time
	password1 := RandomPassword()
	password2 := RandomPassword()
	require.NotEqual(t, password1, password2)
}
