package service

import (
	"context"

	product "bytetech/course/app/frontend/hertz_gen/frontend/product"
	"bytetech/course/app/frontend/infra/rpc"

	"github.com/cloudwego/hertz/pkg/common/utils"

	rpcproduct "bytetech/course/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/app"
)

type GetProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductService(Context context.Context, RequestContext *app.RequestContext) *GetProductService {
	return &GetProductService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductService) Run(req *product.ProductReq) (resp map[string]any, err error) {
	p, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"item": p.Product,
	}, nil
}
