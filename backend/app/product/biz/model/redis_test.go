package model

import (
	"context"
	"errors"
	"log"
	"testing"
	"time"

	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/stretchr/testify/require"
)

func TestRedisEmptyString(t *testing.T) {
	ctx := context.Background()
	testRedisClient.Set(ctx, "test", "", time.Hour)
	res := testRedisClient.Get(ctx, "test")
	log.Printf("空值字符串的字节流长度：%d", len(res.Val()))
}

type TestStruct struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func TestGetValueByKeyFromCache(t *testing.T) {
	key1, value1 := "test", &TestStruct{
		ID:       1,
		Email:    "EMAIL",
		Username: "test",
	}
	data, _ := json.Marshal(value1)
	testRedisClient.Set(context.Background(), key1, data, time.Hour)
	var value2 *TestStruct
	err := getValueByKeyFromCache(context.Background(), testRedisClient, key1, &value2)
	log.Printf("value2: %v", value2)
	log.Printf("err: %v", err)
	require.NotEqual(t, err, cacheEmptyErr)
	require.Equal(t, value1, value2)
}

func TestAAA(t *testing.T) {
	var err error
	require.Nil(t, err)
	require.Equal(t, false, errors.Is(err, cacheEmptyErr))
	err = cacheEmptyErr
	require.Equal(t, true, errors.Is(err, cacheEmptyErr))
}
