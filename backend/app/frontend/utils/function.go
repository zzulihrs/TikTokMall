package utils

import (
	"context"
)

func GetMerchantIdFromCtx(ctx context.Context) int64 {
	// TODO: 需要更改类型信息。目前拿到的 id 使 int64
	userId := ctx.Value(MerchantIdKey)
	if userId == nil {
		return int64(0)
	}
	a, ok := userId.(int64)
	if !ok {
		return int64(0)
	}
	return a
}

func GetUserIdFromCtx(ctx context.Context) uint32 {
	// TODO: 需要更改类型信息。目前拿到的 id 使 int64
	userId := ctx.Value(UserIdKey)
	if userId == nil {
		return uint32(0)
	}
	a, ok := userId.(int32)
	if !ok {
		return uint32(0)
	}
	return uint32(a)
}

func GetUsernameFromCtx(ctx context.Context) string {
	username := ctx.Value(UsernameKey)
	if username == nil {
		return ""
	}
	a, ok := username.(string)
	if !ok {
		return ""
	}
	return a
}
func GetEmailFromCtx(ctx context.Context) string {
	email := ctx.Value(EmailKey)
	if email == nil {
		return ""
	}
	a, ok := email.(string)
	if !ok {
		return ""
	}
	return a
}
