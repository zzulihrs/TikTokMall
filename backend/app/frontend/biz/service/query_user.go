package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	user "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/user"
	"github.com/tiktokmall/backend/app/frontend/infra/rpc"
	user2 "github.com/tiktokmall/backend/rpc_gen/kitex_gen/user"
)

type QueryUserService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewQueryUserService(Context context.Context, RequestContext *app.RequestContext) *QueryUserService {
	return &QueryUserService{RequestContext: RequestContext, Context: Context}
}

func (h *QueryUserService) Run(req *user.QueryUserReq) (resp *user.QueryUserResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	userResp, err := rpc.UserClient.Query(h.Context, &user2.QueryUserReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	resp = &user.QueryUserResp{}
	resp.StatusCode = consts.StatusOK
	resp.Msg = "query successfully"
	resp.User = &user.User{
		Email:    userResp.Email,
		Password: userResp.Password,
	}
	return
}
