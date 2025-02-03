package service

import (
	"context"

	common "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/common"
	"github.com/tiktokmall/backend/app/frontend/infra/rpc"
	"github.com/tiktokmall/backend/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/app"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (resp map[string]any, err error) {
	products, err := rpc.ProductClient.ListProducts(h.Context, &product.ListProductsReq{})
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"title": "Hot Sale",
		"items": products.Products,
	}, nil
}
