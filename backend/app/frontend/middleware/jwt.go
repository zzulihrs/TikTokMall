package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/tiktokmall/backend/app/frontend/biz/utils"
	frontendutils "github.com/tiktokmall/backend/app/frontend/utils"
	"log"
)

func Jwt() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		log.Printf("Jwt")
		tokenString := c.Request.Header.Get("Authorization")

		log.Println("tokenString: ", tokenString)
		if tokenString == "" {
			c.JSON(401, "error request does not contain an access token")
			c.Abort()
			return
		}
		log.Printf("Jwt1")

		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.JSON(401, err.Error())
			c.Abort()
			return
		}

		log.Println("frontendutils.UserIdKey", ctx.Value(frontendutils.UserIdKey))

		//log.Printf("Jwt2")
		log.Println("claims: ", claims.UserId)

		if claims.UserId != 0 {
			ctx = context.WithValue(ctx, frontendutils.UserIdKey, claims.UserId)
		}
		if claims.Username != "" {
			ctx = context.WithValue(ctx, frontendutils.UsernameKey, claims.Username)
		}
		if claims.Email != "" {
			ctx = context.WithValue(ctx, frontendutils.EmailKey, claims.Email)
		}
		if claims.MerchantId != 0 {
			ctx = context.WithValue(ctx, frontendutils.MerchantIdKey, claims.MerchantId)
		}

		c.Next(ctx)
	}
}
