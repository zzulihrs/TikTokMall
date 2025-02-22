package utils

import (
	"context"
)

func GetUserIdFromCtx(ctx context.Context) uint32 {
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
