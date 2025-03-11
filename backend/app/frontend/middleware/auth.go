package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	frontendutils "github.com/tiktokmall/backend/app/frontend/utils"
)

// 全局的身份认证，放到 context 中，方便获取
func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {

		c.Next(ctx)
	}
}

//func GlobalAuth() app.HandlerFunc {
//	return func(ctx context.Context, c *app.RequestContext) {
//		log.Printf("GlobalAuth session.ID: ")
//		tokenString := ctx.Value("Authorization").(string)
//
//		claims, _ := utils.ValidateToken(tokenString)
//
//		if claims.UserId != 0 {
//			ctx = context.WithValue(ctx, frontendutils.UserIdKey, claims.UserId)
//		}
//		if claims.Username != "" {
//			ctx = context.WithValue(ctx, frontendutils.UsernameKey, claims.Username)
//		}
//		if claims.Email != "" {
//			ctx = context.WithValue(ctx, frontendutils.EmailKey, claims.Email)
//		}
//		if claims.MerchantId != 0 {
//			ctx = context.WithValue(ctx, frontendutils.MerchantIdKey, claims.MerchantId)
//		}
//		c.Next(ctx)
//	}
//}

func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		userId := frontendutils.GetUserIdFromCtx(ctx)
		if userId == 0 {
			c.JSON(401, "uid 为空")
			c.Abort()
			return
		}
		email := frontendutils.GetEmailFromCtx(ctx)
		if email == "" {
			c.JSON(401, "email 为空")
			c.Abort()
			return
		}
		c.Next(ctx)
	}
}

func MerchantProduct() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		userId := frontendutils.GetUserIdFromCtx(ctx)
		if userId == 0 {
			c.JSON(401, "uid 为空")
			c.Abort()
			return
		}
		// log.Printf("userId: %v", userId)
		email := frontendutils.GetEmailFromCtx(ctx)
		if email == "" {
			c.JSON(401, "email 为空")
			c.Abort()
			return
		}
		// log.Printf("email: %v", email)
		merchantId := frontendutils.GetMerchantIdFromCtx(ctx)
		if merchantId == 0 {
			c.JSON(403, "merchant_id 为空")
			c.Abort()
			return
		}
		// log.Printf("merchantId: %v", merchantId)
		c.Next(ctx)

	}
}
