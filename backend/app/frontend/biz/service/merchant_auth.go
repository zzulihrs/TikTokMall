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

/*
Request
{
	"uid": 100,
    "email": "cloudwego.com"
}
*/

func (h *MerchantAuthService) Run(req *merchant.MerchantAuthReq) (resp utils.H, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	// 1. 调用 merchantrpc 获取店家信息
	// 2. 包装店家信息生成 token
	// 3. 返回 token 和店家信息 response
	return
}

/*
Response
{
    "code": 61,
    "message": "nulla reprehenderit proident",
    "token": "aute consectetur Ut",
    "merchant_info": {
        "id": 98,
        "username": "戊国良",
        "email": "nlrlqi.px58@gmail.com"
    }
}
*/
