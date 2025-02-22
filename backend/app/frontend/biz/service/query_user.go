package service

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	user "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/user"
	"github.com/tiktokmall/backend/app/frontend/infra/rpc"
	frontendUtils "github.com/tiktokmall/backend/app/frontend/utils"
	rpcuser "github.com/tiktokmall/backend/rpc_gen/kitex_gen/user"
)

type QueryUserService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewQueryUserService(Context context.Context, RequestContext *app.RequestContext) *QueryUserService {
	return &QueryUserService{RequestContext: RequestContext, Context: Context}
}

func (h *QueryUserService) Run(req *user.QueryUserReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	uid := frontendUtils.GetUserIdFromCtx(h.Context)
	username := frontendUtils.GetUsernameFromCtx(h.Context)
	email := frontendUtils.GetEmailFromCtx(h.Context)
	if uid == 0 {
		return nil, fmt.Errorf("用户未登录")
	}
	if username != "" && email != "" {
		return map[string]any{
			"code":    200,
			"message": "ok",
			"data": map[string]any{
				"id":       uid,
				"username": username,
				"email":    email,
			},
		}, nil
	}
	// 数据库
	userResp, err := rpc.UserClient.Query(h.Context, &rpcuser.QueryUserReq{
		UserId: int64(uid),
	})
	if err != nil {
		return nil, err
	}
	resp = map[string]any{
		"code":    200,
		"message": "ok",
		"data": map[string]any{
			"id":       uid,
			"username": userResp.User.Username,
			"email":    userResp.User.Email,
		},
	}
	return resp, nil
}
