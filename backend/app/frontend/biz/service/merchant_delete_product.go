package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	merchant "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/merchant"
	"github.com/tiktokmall/backend/app/frontend/infra/rpc"
	frontendUtils "github.com/tiktokmall/backend/app/frontend/utils"
	rpcmerchant "github.com/tiktokmall/backend/rpc_gen/kitex_gen/merchant"
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
	_, err = rpc.MerchantClient.DeleteProduct(h.Context, &rpcmerchant.DeleteProductReq{
		MerchantId: frontendUtils.GetMerchantIdFromCtx(h.Context),
		Pid:        req.Pid,
	})
	if err != nil {
		return utils.H{
			"code":    500,
			"message": "删除商品失败",
		}, err
	}
	return utils.H{
		"code":    200,
		"message": "OK",
	}, nil
}

/*
Response
{
    "code": 200,
    "message": "OK"
}
*/
