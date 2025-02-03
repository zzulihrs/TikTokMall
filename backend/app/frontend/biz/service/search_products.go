package service

import (
	"context"

	product "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/product"
	"github.com/tiktokmall/backend/app/frontend/infra/rpc"
	rpcproduct "github.com/tiktokmall/backend/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/app"
)

type SearchProductsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProductsService(Context context.Context, RequestContext *app.RequestContext) *SearchProductsService {
	return &SearchProductsService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProductsService) Run(req *product.SearchProductsReq) (resp map[string]any, err error) {
	products, err := rpc.ProductClient.SearchProducts(h.Context, &rpcproduct.SearchProductsReq{
		Query: req.Q,
	})
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"items": products.Results,
		"q":     req.Q,
	}, nil
}
