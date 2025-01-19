package utils

import (
	"context"
)

func GetUserIdFromCtx(ctx context.Context) uint32 {
	userId := ctx.Value(UserIdKey)
	if userId == nil {
		return uint32(0)
	}
	// fmt.Printf("UserId: %v, Type: %T\n", userId, userId)
	a, ok := userId.(int32)
	if !ok {
		// fmt.Println("Error: UserId is not an int32")
		return uint32(0)
	}
	return uint32(a)
}
