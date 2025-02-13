package service

import (
	"context"
	"github.com/tiktokmall/backend/app/frontend/infra/rpc"
	frontendUtils "github.com/tiktokmall/backend/app/frontend/utils"
	rpccart "github.com/tiktokmall/backend/rpc_gen/kitex_gen/cart"

	"github.com/cloudwego/hertz/pkg/app"
	cart "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/cart"
	common "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/common"
)

type ChangeQtyService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewChangeQtyService(Context context.Context, RequestContext *app.RequestContext) *ChangeQtyService {
	return &ChangeQtyService{RequestContext: RequestContext, Context: Context}
}

func (h *ChangeQtyService) Run(req *cart.ChangeQtyReq) (resp *common.Empty, err error) {
	_, err = rpc.CartClient.ChangeQty(h.Context, &rpccart.ChangeQtyReq{
		UserId: uint32(frontendUtils.GetUserIdFromCtx(h.Context)),
		Item: &rpccart.CartItem{
			ProductId: req.GetProductId(),
			Quantity:  uint32(req.GetProductNum()),
		},
	})
	if err != nil {
		return nil, err
	}

	return
}
