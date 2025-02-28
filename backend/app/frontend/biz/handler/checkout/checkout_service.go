package checkout

import (
	"context"

	"github.com/tiktokmall/backend/app/frontend/biz/service"
	"github.com/tiktokmall/backend/app/frontend/biz/utils"
	checkout "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/checkout"
	common "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/common"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Checkout .
// @router /checkout [GET]
func Checkout(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]any{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	resp, err := service.NewCheckoutService(ctx, c).Run(&req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, map[string]any{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	resp["code"] = 200
	resp["message"] = "ok"
	c.JSON(consts.StatusOK, utils.WarpResponse(ctx, c, resp))
}

// CheckoutWaiting .
// @router /checkout/waiting [POST]
func CheckoutWaiting(ctx context.Context, c *app.RequestContext) {
	var err error
	var req checkout.CheckoutReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]any{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	resp, err := service.NewCheckoutWaitingService(ctx, c).Run(&req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, map[string]any{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	resp["code"] = 200
	resp["message"] = "ok"
	c.JSON(consts.StatusOK, utils.WarpResponse(ctx, c, resp))
}

// CheckoutResult .
// @router /checkout/result [GET]
func CheckoutResult(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]any{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	resp, err := service.NewCheckoutResultService(ctx, c).Run(&req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, map[string]any{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	resp["code"] = 200
	resp["message"] = "ok"
	c.JSON(consts.StatusOK, utils.WarpResponse(ctx, c, resp))
}
