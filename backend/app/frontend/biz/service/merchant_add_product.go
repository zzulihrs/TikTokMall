package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	merchant "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/merchant"
	"github.com/tiktokmall/backend/app/frontend/infra/rpc"
	rpcmerchant "github.com/tiktokmall/backend/rpc_gen/kitex_gen/merchant"
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
	categories := make([]*rpcmerchant.Category, len(req.Categories))
	for i, category := range req.Categories {
		categories[i] = &rpcmerchant.Category{
			Id:          category.Id,
			Name:        category.Name,
			Description: category.Description,
		}
	}
	_, err = rpc.MerchantClient.AddProduct(h.Context, &rpcmerchant.AddProductReq{
		MerchantId: req.Mid,
		Product: &rpcmerchant.MerchantProductDetailInfo{
			Name:        req.Name,
			Description: req.Description,
			Stock:       req.Stock,
			Price:       req.Price,
			ImgUrl:      req.ImgUrl,
			SliderImgs:  req.SliderImgs,
			Category:    categories,
		},
	})
	if err != nil {
		return utils.H{
			"code":    500,
			"message": "添加商品失败",
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
