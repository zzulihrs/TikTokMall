package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	merchant "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/merchant"
	"github.com/tiktokmall/backend/app/frontend/utils"
)

type MerchantGetProductListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewMerchantGetProductListService(Context context.Context, RequestContext *app.RequestContext) *MerchantGetProductListService {
	return &MerchantGetProductListService{RequestContext: RequestContext, Context: Context}
}

func (h *MerchantGetProductListService) Run(req *merchant.MerchantGetProductListReq) (resp utils.H, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
