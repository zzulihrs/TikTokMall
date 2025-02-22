package service

import (
	"context"

	auth "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/auth"
	"github.com/tiktokmall/backend/app/frontend/infra/rpc"
	"github.com/tiktokmall/backend/rpc_gen/kitex_gen/user"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *auth.RegisterReq) (user_id int64, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// TODO: user svc api

	userResp, err := rpc.UserClient.Register(h.Context, &user.RegisterReq{
		Email:           req.Email,
		Password:        req.Password,
		ConfirmPassword: req.ConfirmPassword,
	})
	if err != nil {
		return 0, err
	}

	session := sessions.Default(h.RequestContext)

	session.Set("user_id", userResp.UserId)
	err = session.Save()
	if err != nil {
		return 0, err
	}
	return int64(userResp.UserId), nil
}
