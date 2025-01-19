package middleware

import (
	frontendutils "bytetech/course/app/frontend/utils"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

// 全局的身份认证，放到 context 中，方便获取
func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		userId := session.Get("user_id")
		if userId == nil {
			c.Next(ctx)
			return
		}
		ctx = context.WithValue(ctx, frontendutils.UserIdKey, userId)
		c.Next(ctx)

	}
}

func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		userId := session.Get("user_id")

		if userId == nil {
			c.Redirect(302, []byte("/sign-in?next="+c.FullPath()))
			c.Abort()
			return
		}
		// ctx = context.WithValue(ctx, frontendutils.UserIdKey, userId)
		c.Next(ctx)

	}
}
