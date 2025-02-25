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
	pResp, err := rpc.MerchantClient.ProductDetail(h.Context, &rpcmerchant.ProductDetailReq{
		MerchantId: frontendUtils.GetMerchantIdFromCtx(h.Context),
		Pid:        req.Id,
	})
	if err != nil {
		return utils.H{
			"code":    500,
			"message": "获取商品详情失败",
		}, err
	}
	categories := make([]utils.H, len(pResp.Product.Category))
	for i, c := range pResp.Product.Category {
		categories[i] = utils.H{
			"id":          c.Id,
			"name":        c.Name,
			"description": c.Description,
		}
	}
	product := utils.H{
		"id":          pResp.Product.Id,
		"name":        pResp.Product.Name,
		"description": pResp.Product.Description,
		"stock":       pResp.Product.Stock,
		"price":       pResp.Product.Price,
		"img_url":     pResp.Product.ImgUrl,
		"category":    categories,
		"slider_img":  pResp.Product.SliderImgs,
	}
	return utils.H{
		"code":    200,
		"message": "OK",
		"product": product,
	}, nil
}

/*
Response
{
    "code": 200,
    "message": "sss",
    "product": {
        "id": 1,
        "name": "aa",
        "description": "aaa",
        "stock": 23,
        "price": 19.26,
        "img_url": "static/img.png",
        "catetory": [
            {
                "id": 1,
                "name": "t-shirt",
                "description": "t-shirt"
            }
        ],
        "slider_img": [
            "static/img.png"
        ]
    }
}
*/
