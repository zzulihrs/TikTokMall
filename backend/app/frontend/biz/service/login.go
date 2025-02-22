package service

import (
	"context"

	auth "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/auth"
	"github.com/tiktokmall/backend/app/frontend/infra/rpc"
	"github.com/tiktokmall/backend/rpc_gen/kitex_gen/user"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	rpc_resp, err := rpc.UserClient.Login(h.Context, &user.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})
	// 没有登陆成功
	if err != nil {
		return nil, err
	}
	// 获取 请求的上下文 得到 session
	session := sessions.Default(h.RequestContext)
	// 利用 redis 存储 user_id
	// 设置 Cookie
	// Cookie

	session.Set("user_id", rpc_resp.UserId)
	session.Set("username", rpc_resp.Username)
	session.Set("email", rpc_resp.Email)
	err = session.Save()
	if err != nil {
		return nil, err
	}
	// 从 request 中获取 referer ，然后返回上一个 url
	// redirect := "/"
	// if req.Next != "" {
	// 	redirect = req.Next
	// }
	return map[string]any{
		"status":  200,
		"message": "OK",
	}, err
}
