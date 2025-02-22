package middleware

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
	"github.com/joho/godotenv"
	"github.com/tiktokmall/backend/app/frontend/biz/service"
	"github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/auth"
	u "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/user"
)

var JwtMiddleware *jwt.HertzJWTMiddleware

func InitJWT() {
	_ = godotenv.Load()
	JwtMiddleware, _ = jwt.New(&jwt.HertzJWTMiddleware{
		Key:           []byte(os.Getenv("SecretKey")),
		TokenLookup:   "form: token, param: token, header: Authorization, query: token", //这里主要用到form和query
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		Timeout:       time.Hour,                //指定了 token 有效期为一个小时
		MaxRefresh:    time.Hour,                //用于设置最大 token 刷新时间
		IdentityKey:   os.Getenv("IdentityKey"), // 用于设置检索身份的键
		////用于设置登录验证的函数----专职服务login
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var err error
			var req auth.LoginReq
			if err = c.BindAndValidate(&req); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			if len(req.Email) == 0 || len(req.Password) == 0 {
				return "", jwt.ErrMissingLoginValues
			}
			//调用rpc方法
			resp, err := service.NewLoginService(ctx, c).Run(&req)
			c.Set(os.Getenv("IdentityKey"), resp["id"]) //将ID存到上下文中，然后返回报文时再取
			return resp, err
		},
		//用于设置获取身份信息的函数，此处提取 token 的负载，并配合 IdentityKey 将用户id存入上下文信息。
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &u.User{
				UserId: int64(claims[os.Getenv("IdentityKey")].(float64)),
			}
		},
		//它的入参就是 Authenticator 的返回值，此时负责解析 user，并将用户id注入 token 的 payload 部分
		//这个函数将配合IdentityHandler
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(map[string]any); ok {
				return jwt.MapClaims{
					os.Getenv("IdentityKey"): v["id"],
				}
			}
			return jwt.MapClaims{}
		},
		//用于设置登录的响应函数----专职服务login
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			v, _ := c.Get("id")
			c.JSON(code, utils.H{
				"status_code": 0,
				"status_msg":  "success",
				"user_id":     v,
				"token":       token,
			})
		},

		/*下面是登录失败的响应函数*/
		//用于设置 jwt 验证流程失败的响应函数，当前 demo 返回了错误码和错误信息。
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, utils.H{
				"status_code": 10003,
				"status_msg":  message,
			})
		},
		//用于设置 jwt 校验流程发生错误时响应所包含的错误信息，你可以自行包装这些内容
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			return e.Error()
		},
	})
}
