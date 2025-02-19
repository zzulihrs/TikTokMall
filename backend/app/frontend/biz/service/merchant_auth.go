package service

import (
	"context"
	"log"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/tiktokmall/backend/app/frontend/biz/utils"
	merchant "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/merchant"
	"github.com/tiktokmall/backend/app/frontend/infra/rpc"

	rpcmerchant "github.com/tiktokmall/backend/rpc_gen/kitex_gen/merchant"
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
	log.Printf("MerchantAuthService Run, req = %+v", req)
	merchantResp, err := rpc.MerchantClient.GetMerchant(h.Context, &rpcmerchant.GetMerchantReq{
		Id: req.Uid,
	})
	if err != nil {
		return nil, err
	}
	// 2. 包装店家信息生成 token
	token := utils.GenerateToken(merchantResp.Id, merchantResp.Username)
	// 3. 返回 token 和店家信息 response
	resp = utils.H{
		"code":    200,
		"message": "OK",
		"token":   token,
		"merchant_info": utils.H{
			"id":       merchantResp.Id,
			"username": merchantResp.Username,
			"email":    merchantResp.Email,
		},
	}
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
