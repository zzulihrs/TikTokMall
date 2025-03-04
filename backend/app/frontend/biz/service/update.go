package service

import (
	"context"

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

func (h *UpdateService) Run(req *user2.UpdateUserReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	_, err = rpc.UserClient.Update(h.Context, &user.UpdateUserReq{
		Username: req.Username,
		Avator:   req.Avator,
		UserId:   int64(frontendUtils.GetUserIdFromCtx(h.Context)),
		Email:    frontendUtils.GetEmailFromCtx(h.Context),
	})
	if err != nil {
		return nil, err
	}

	resp = map[string]any{
		"code":    200,
		"message": "ok",
	}
	return
}
