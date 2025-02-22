package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	merchant "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/merchant"
	"github.com/tiktokmall/backend/app/frontend/infra/rpc"
	frontendUtils "github.com/tiktokmall/backend/app/frontend/utils"
	rpcmerchant "github.com/tiktokmall/backend/rpc_gen/kitex_gen/merchant"
)

type MerchantRegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewMerchantRegisterService(Context context.Context, RequestContext *app.RequestContext) *MerchantRegisterService {
	return &MerchantRegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *MerchantRegisterService) Run(req *merchant.MerchantRegisterReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	if req.Code != "Hello World" {
		return map[string]any{
			"code":    400,
			"message": "验证码错误",
		}, nil
	}

	// NOTE: 店家的名称，默认采用 user 的名称
	// TODO：后续可以考虑进一步优化表结构。
	userId := frontendUtils.GetUserIdFromCtx(h.Context)
	username := frontendUtils.GetUsernameFromCtx(h.Context)
	email := frontendUtils.GetEmailFromCtx(h.Context)
	rpc.MerchantClient.AddMerchant(h.Context, &rpcmerchant.AddMerchantReq{
		Uid:      int64(userId),
		Username: username,
		Email:    email,
	})

	return map[string]any{
		"code":    200,
		"message": "ok",
	}, nil
}
