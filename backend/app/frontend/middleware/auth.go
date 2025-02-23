package middleware

import (
	"context"

	frontendutils "github.com/tiktokmall/backend/app/frontend/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

// 全局的身份认证，放到 context 中，方便获取
func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		userId := session.Get("user_id")
		username := session.Get("username")
		email := session.Get("email")
		if userId == nil {
			c.Next(ctx)
			return
		}
		ctx = context.WithValue(ctx, frontendutils.UserIdKey, userId)
		ctx = context.WithValue(ctx, frontendutils.UsernameKey, username)
		ctx = context.WithValue(ctx, frontendutils.EmailKey, email)
		c.Next(ctx)

	}
}

func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		userId := session.Get("user_id")
		username := session.Get("username")
		email := session.Get("email")
		merchantId := session.Get("merchant_id")
		if userId == nil {
			c.JSON(302, "uid 为空")
			// c.Redirect(302, []byte("/sign-in?next="+c.FullPath()))
			c.Abort()
			return
		}
		ctx = context.WithValue(ctx, frontendutils.UserIdKey, userId)
		ctx = context.WithValue(ctx, frontendutils.UsernameKey, username)
		ctx = context.WithValue(ctx, frontendutils.EmailKey, email)
		ctx = context.WithValue(ctx, frontendutils.MerchantIdKey, merchantId)
		c.Next(ctx)

	}
}

func MerchantProduct() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		userId := session.Get("user_id")
		username := session.Get("username")
		email := session.Get("email")
		if userId == nil {
			c.JSON(302, "uid 为空")
			c.Abort()
			return
		}
		merchantId := session.Get("merchant_id")
		// log.Printf("userId: %v", userId)
		// log.Printf("merchantId: %v", merchantId)
		// log.Printf("username: %v", username)
		// log.Printf("email: %v", email)
		if merchantId == nil {
			c.JSON(302, "merchant_id 为空")
			c.Abort()
			return
		}
		ctx = context.WithValue(ctx, frontendutils.UserIdKey, userId)
		ctx = context.WithValue(ctx, frontendutils.UsernameKey, username)
		ctx = context.WithValue(ctx, frontendutils.EmailKey, email)
		ctx = context.WithValue(ctx, frontendutils.MerchantIdKey, merchantId)

		c.Next(ctx)

	}
}
