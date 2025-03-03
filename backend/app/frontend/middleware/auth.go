package middleware

import (
	"context"

	"github.com/tiktokmall/backend/app/frontend/utils"
	frontendutils "github.com/tiktokmall/backend/app/frontend/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

// 全局的身份认证，放到 context 中，方便获取
func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		// log.Printf("GlobalAuth session.ID: %v", session.ID())
		userId := session.Get("user_id")
		if userId != nil {
			ctx = context.WithValue(ctx, frontendutils.UserIdKey, userId)
		}
		username := session.Get("username")
		if username != nil {
			ctx = context.WithValue(ctx, frontendutils.UsernameKey, username)
		}
		email := session.Get("email")
		if email != nil {
			ctx = context.WithValue(ctx, frontendutils.EmailKey, email)
		}
		merchantId := session.Get("merchant_id")
		if merchantId != nil {
			ctx = context.WithValue(ctx, frontendutils.MerchantIdKey, merchantId)
		}
		c.Next(ctx)
	}
}

func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		userId := utils.GetUserIdFromCtx(ctx)
		if userId == 0 {
			c.JSON(302, "uid 为空")
			c.Abort()
			return
		}
		email := utils.GetEmailFromCtx(ctx)
		if email == "" {
			c.JSON(302, "email 为空")
			c.Abort()
			return
		}
		c.Next(ctx)
	}
}

func MerchantProduct() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		userId := utils.GetUserIdFromCtx(ctx)
		if userId == 0 {
			c.JSON(302, "uid 为空")
			c.Abort()
			return
		}
		// log.Printf("userId: %v", userId)
		email := utils.GetEmailFromCtx(ctx)
		if email == "" {
			c.JSON(302, "email 为空")
			c.Abort()
			return
		}
		// log.Printf("email: %v", email)
		merchantId := utils.GetMerchantIdFromCtx(ctx)
		if merchantId == 0 {
			c.JSON(302, "merchant_id 为空")
			c.Abort()
			return
		}
		// log.Printf("merchantId: %v", merchantId)
		c.Next(ctx)

	}
}
