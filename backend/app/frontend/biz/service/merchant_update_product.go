package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	merchant "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/merchant"
)

type MerchantUpdateProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewMerchantUpdateProductService(Context context.Context, RequestContext *app.RequestContext) *MerchantUpdateProductService {
	return &MerchantUpdateProductService{RequestContext: RequestContext, Context: Context}
}

func (h *MerchantUpdateProductService) Run(req *merchant.MerchantUpdateProductReq) (resp utils.H, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
