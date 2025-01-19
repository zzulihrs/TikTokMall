package service

import (
	"context"

	cart "bytetech/course/app/frontend/hertz_gen/frontend/cart"
	"bytetech/course/app/frontend/infra/rpc"
	frontendUtils "bytetech/course/app/frontend/utils"
	rpccart "bytetech/course/rpc_gen/kitex_gen/cart"

	"github.com/cloudwego/hertz/pkg/app"
)

type AddCartItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddCartItemService(Context context.Context, RequestContext *app.RequestContext) *AddCartItemService {
	return &AddCartItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddCartItemService) Run(req *cart.AddCartReq) (resp map[string]any, err error) {
	_, err = rpc.CartClient.AddItem(h.Context, &rpccart.AddItemReq{
		UserId: frontendUtils.GetUserIdFromCtx(h.Context),
		Item: &rpccart.CartItem{
			ProductId: req.GetProductId(),
			Quantity:  req.GetProductNum(),
		},
	})
	return map[string]any{}, err
}
