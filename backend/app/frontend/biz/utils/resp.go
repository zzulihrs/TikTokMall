package utils

import (
	"context"

	"github.com/tiktokmall/backend/app/frontend/infra/rpc"
	frontendutils "github.com/tiktokmall/backend/app/frontend/utils"
	rpccart "github.com/tiktokmall/backend/rpc_gen/kitex_gen/cart"

	"github.com/cloudwego/hertz/pkg/app"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	c.String(code, err.Error())
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, data)
}

// SendSuccessResponse  pack success response
func WarpResponse(ctx context.Context, c *app.RequestContext, resp map[string]any) map[string]any {
	if resp == nil {
		resp = map[string]any{}
	}
	userId := frontendutils.GetUserIdFromCtx(ctx)
	resp["user_id"] = userId
	cartResp, _ := rpc.CartClient.GetCart(ctx, &rpccart.GetCartReq{UserId: userId})
	var cartNum int
	if cartResp != nil && cartResp.Cart != nil {
		cartNum = len(cartResp.Cart.Items)
	}
	resp["cart_num"] = cartNum
	return resp
}
