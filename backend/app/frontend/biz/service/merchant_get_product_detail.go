package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	merchant "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/merchant"
)

type MerchantGetProductDetailService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewMerchantGetProductDetailService(Context context.Context, RequestContext *app.RequestContext) *MerchantGetProductDetailService {
	return &MerchantGetProductDetailService{RequestContext: RequestContext, Context: Context}
}

func (h *MerchantGetProductDetailService) Run(req *merchant.MerchantGetProductDetailReq) (resp utils.H, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
