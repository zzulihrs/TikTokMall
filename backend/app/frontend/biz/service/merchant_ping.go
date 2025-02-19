package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	common "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/common"
	"github.com/tiktokmall/backend/app/frontend/utils"
)

type MerchantPingService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewMerchantPingService(Context context.Context, RequestContext *app.RequestContext) *MerchantPingService {
	return &MerchantPingService{RequestContext: RequestContext, Context: Context}
}

func (h *MerchantPingService) Run(req *common.Empty) (resp utils.H, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
