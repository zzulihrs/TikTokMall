package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	merchant "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/merchant"
	"github.com/tiktokmall/backend/app/frontend/utils"
)

type MerchantAuthService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewMerchantAuthService(Context context.Context, RequestContext *app.RequestContext) *MerchantAuthService {
	return &MerchantAuthService{RequestContext: RequestContext, Context: Context}
}

func (h *MerchantAuthService) Run(req *merchant.MerchantAuthReq) (resp utils.H, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
