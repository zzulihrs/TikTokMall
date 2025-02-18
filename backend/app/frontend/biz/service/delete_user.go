package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/sessions"
	user "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/user"
	"github.com/tiktokmall/backend/app/frontend/infra/rpc"
	user2 "github.com/tiktokmall/backend/rpc_gen/kitex_gen/user"
)

type DeleteUserService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteUserService(Context context.Context, RequestContext *app.RequestContext) *DeleteUserService {
	return &DeleteUserService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteUserService) Run(req *user.DeleteUserReq) (resp *user.DeleteUserResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	_, err = rpc.UserClient.Delete(h.Context, &user2.DeleteUserReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	session := sessions.Default(h.RequestContext)
	session.Clear()
	err = session.Save()
	if err != nil {
		return nil, err
	}
	resp = &user.DeleteUserResp{}
	resp.StatusCode = consts.StatusOK
	resp.Msg = "delete successfully"
	return resp, nil
}
