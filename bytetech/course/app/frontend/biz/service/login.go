package service

import (
	"context"

	auth "bytetech/course/app/frontend/hertz_gen/frontend/auth"
	"bytetech/course/app/frontend/infra/rpc"
	"bytetech/course/rpc_gen/kitex_gen/user"

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

func (h *LoginService) Run(req *auth.LoginReq) (redirect string, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	resp, err := rpc.UserClient.Login(h.Context, &user.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})
	// 没有登陆成功
	if err != nil {
		return "", err
	}
	// 获取 请求的上下文 得到 session
	session := sessions.Default(h.RequestContext)
	// 利用 redis 存储 user_id
	session.Set("user_id", resp.UserId)
	err = session.Save()
	if err != nil {
		return "", err
	}
	// 从 request 中获取 referer ，然后返回上一个 url
	redirect = "/"
	if req.Next != "" {
		redirect = req.Next
	}
	return
}
