package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	merchant "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/merchant"
)

type MerchantAddProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewMerchantAddProductService(Context context.Context, RequestContext *app.RequestContext) *MerchantAddProductService {
	return &MerchantAddProductService{RequestContext: RequestContext, Context: Context}
}

func (h *MerchantAddProductService) Run(req *merchant.MerchantAddProductReq) (resp utils.H, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}

/*
Response
{
    "code": 58,
    "message": "pariatur deserunt ullamco labore"
}
*/
