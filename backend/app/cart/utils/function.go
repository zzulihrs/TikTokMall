package utils

import (
	"context"
	"fmt"
)

func GetUserIdFromCtx(ctx context.Context) int32 {
	userId := ctx.Value(SessionUserId)
	if userId == nil {
		return 0
	}
	fmt.Println("userId", userId)

	return userId.(int32)
}
