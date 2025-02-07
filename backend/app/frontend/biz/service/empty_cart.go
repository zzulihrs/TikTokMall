package service

import (
	"context"
	"github.com/tiktokmall/backend/app/frontend/infra/rpc"
	frontendUtils "github.com/tiktokmall/backend/app/frontend/utils"
	rpccart "github.com/tiktokmall/backend/rpc_gen/kitex_gen/cart"

	"github.com/cloudwego/hertz/pkg/app"
	common "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/common"
)

type EmptyCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewEmptyCartService(Context context.Context, RequestContext *app.RequestContext) *EmptyCartService {
	return &EmptyCartService{RequestContext: RequestContext, Context: Context}
}

func (h *EmptyCartService) Run(req *common.Empty) (resp map[string]any, err error) {
	_, err = rpc.CartClient.EmptyCart(h.Context, &rpccart.EmptyCartReq{
		UserId: frontendUtils.GetUserIdFromCtx(h.Context),
	})
	return map[string]any{
		"message": "success",
	}, err
}
