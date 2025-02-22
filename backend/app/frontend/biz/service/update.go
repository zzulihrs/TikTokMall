package service

import (
	"context"
	"fmt"
	frontendUtils "github.com/tiktokmall/backend/app/frontend/utils"

	"github.com/cloudwego/hertz/pkg/app"
	user2 "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/user"
	"github.com/tiktokmall/backend/app/frontend/infra/rpc"
	"github.com/tiktokmall/backend/rpc_gen/kitex_gen/user"
)

type UpdateService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateService(Context context.Context, RequestContext *app.RequestContext) *UpdateService {
	return &UpdateService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateService) Run(req *user2.UpdateUserReq) (redirect string, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	resp, err := rpc.UserClient.Update(h.Context, &user.UpdateUserReq{
		UserId: req.UserId,
		Email:  req.Email,
		Avatar: req.Avatar,
	})
	// 没有更新成功
	if err != nil {
		return "", err
	}
	redirect = "/:" + fmt.Sprintf("%d", resp.UserId)
	return redirect, nil
}
