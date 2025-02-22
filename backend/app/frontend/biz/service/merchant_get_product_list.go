package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	merchant "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/merchant"
	"github.com/tiktokmall/backend/app/frontend/infra/rpc"
	"github.com/tiktokmall/backend/app/frontend/utils"
	rpcmerchant "github.com/tiktokmall/backend/rpc_gen/kitex_gen/merchant"
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
	// log.Printf("merchant get product list, req: %+v", req)
	psResp, err := rpc.MerchantClient.SearchProducts(h.Context, &rpcmerchant.SearchProductsReq{
		Name:       req.Name,
		CategoryId: req.CategoryId,
		MaxStock:   int32(req.MaxStock),
		MaxPrice:   req.MaxPrice,
		MinPrice:   req.MinPrice,
		PageNo:     req.PageNo,
		PageSize:   req.PageSize,
		MerchantId: req.MerchantId,
	})
	if err != nil {
		return utils.H{
			"code":    500,
			"message": "获取商品列表失败",
		}, err
	}
	products := make([]utils.H, len(psResp.Products))
	for i, p := range psResp.Products {
		products[i] = utils.H{
			"id":          p.Id,
			"name":        p.Name,
			"description": p.Description,
			"stock":       p.Stock,
			"price":       p.Price,
			"img_url":     p.ImgUrl,
		}
	}
	return utils.H{
		"code":    200,
		"message": "OK",
		"data":    products,
		"total":   psResp.Count,
	}, nil
}

/*
Response
{
    "code": 200,
    "message": "proident anim id culpa",
    "data": [
        {
            "id": 85,
            "name": "纳喇超",
            "description": "等于或公万整小。设步当采多争。保更们真效务目那。入都今求会活还展。走角再县走有务内少更。",
            "stock": 83,
            "price": "531.38",
            "img_url": "https://loremflickr.com/400/400?lock=6729080248750675"
        },
        {
            "id": 40,
            "name": "寿婷婷",
            "description": "发会写着产件阶场人。情众着声。选给片走。还九值素写增己如。就七照济。物数你区然气族越。边程育公。",
            "stock": 5,
            "price": "801.29",
            "img_url": "https://loremflickr.com/400/400?lock=3451811131915255"
        },
        {
            "id": 98,
            "name": "伏国英",
            "description": "气往温去。火制当角去观每来。",
            "stock": 43,
            "price": "468.14",
            "img_url": "https://loremflickr.com/400/400?lock=8796401005629023"
        }
    ]
}
*/
