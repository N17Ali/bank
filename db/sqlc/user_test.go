package db

import (
	"context"
	"testing"
	"time"

	"github.com/n17ali/bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}
func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	gotUser, err := testQueries.GetUsers(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, gotUser)

	require.Equal(t, user1.Username, gotUser.Username)
	require.Equal(t, user1.HashedPassword, gotUser.HashedPassword)
	require.Equal(t, user1.FullName, gotUser.FullName)
	require.Equal(t, user1.Email, gotUser.Email)
	require.WithinDuration(t, user1.PasswordChangedAt, gotUser.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, gotUser.CreatedAt, time.Second)
}
