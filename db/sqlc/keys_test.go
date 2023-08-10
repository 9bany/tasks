package db

import (
	"context"
	"testing"

	"github.com/9bany/task/util"
	"github.com/stretchr/testify/require"
)

func createRandomKey(t *testing.T) Keys {

	argKey := util.RandomString(32)

	key, err := testQueries.CreateKey(context.Background(), argKey)

	require.NoError(t, err)
	require.NotEmpty(t, key)

	require.Equal(t, argKey, key.Key)
	require.Equal(t, int32(0), key.UsageCount)

	require.NotZero(t, key.CreatedAt)

	return key
}

func TestCreateKey(t *testing.T) {
	createRandomKey(t)
}

func TestGetRandomKey(t *testing.T) {
	key, err := testQueries.GetRandomKey(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, key)
	require.NotZero(t, key.CreatedAt)
}

func TestIncreaseUsageCount(t *testing.T) {
	key1 := createRandomKey(t)
	err := testQueries.IncreaseKeyUsageCount(context.Background(), key1.ID)
	require.NoError(t, err)
}

func TestGetKeyInfo(t *testing.T) {
	key1 := createRandomKey(t)
	key2, err := testQueries.GetKey(context.Background(), key1.Key)
	
	require.NoError(t, err)
	require.NotEmpty(t, key2)

	require.Equal(t, key1.Key, key2.Key)
}
