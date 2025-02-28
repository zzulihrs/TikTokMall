package cart

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	hertzUtils "github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/tiktokmall/backend/app/frontend/biz/service"
	"github.com/tiktokmall/backend/app/frontend/biz/utils"
	cart "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/cart"
	common "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/common"
)

// AddCartItem .
// @router /cart [POST]
func AddCartItem(ctx context.Context, c *app.RequestContext) {
	var err error
	var req cart.AddCartItemReq
	err = c.BindAndValidate(&req)

	// 打印解析后的数据（用于调试）
	fmt.Println("Parsed Request Data:", req.GetProductId(), req.GetProductNum())

	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]any{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	_, err = service.NewAddCartItemService(ctx, c).Run(&req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, map[string]any{
			"code":    500,
			"message": err.Error(),
		})
		return
	}
	c.JSON(consts.StatusOK, utils.WarpResponse(ctx, c, hertzUtils.H{
		"code":    200,
		"message": "ok",
	}))
}

// GetCart .
// @router /cart [GET]
func GetCart(ctx context.Context, c *app.RequestContext) {
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
	resp, err := service.NewGetCartService(ctx, c).Run(&req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, map[string]any{
			"code":    500,
			"message": err.Error(),
		})
		return
	}
	resp["code"] = 200
	resp["message"] = "OK"
	c.JSON(consts.StatusOK, utils.WarpResponse(ctx, c, resp))
}

// EmptyCart .
// @router /emptyCart [GET]
func EmptyCart(ctx context.Context, c *app.RequestContext) {
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

	resp, err := service.NewEmptyCartService(ctx, c).Run(&req)

	if err != nil {
		c.JSON(consts.StatusInternalServerError, map[string]any{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	resp["code"] = 200
	resp["message"] = "OK"
	c.JSON(consts.StatusOK, resp)
}

// ChangeQty .
// @router /changeqty [POST]
func ChangeQty(ctx context.Context, c *app.RequestContext) {
	var err error
	var req cart.ChangeQtyReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]any{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	_, err = service.NewChangeQtyService(ctx, c).Run(&req)

	if err != nil {
		c.JSON(consts.StatusInternalServerError, map[string]any{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	c.JSON(consts.StatusOK, utils.WarpResponse(ctx, c, hertzUtils.H{
		"code":    200,
		"message": "ok",
	}))

}
