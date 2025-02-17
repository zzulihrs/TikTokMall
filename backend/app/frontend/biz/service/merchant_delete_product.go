package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	merchant "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/merchant"
)

type MerchantDeleteProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewMerchantDeleteProductService(Context context.Context, RequestContext *app.RequestContext) *MerchantDeleteProductService {
	return &MerchantDeleteProductService{RequestContext: RequestContext, Context: Context}
}

func (h *MerchantDeleteProductService) Run(req *merchant.MerchantDeleteProductReq) (resp utils.H, err error) {
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
    "code": 200,
    "message": "delete sucess"
}
*/
