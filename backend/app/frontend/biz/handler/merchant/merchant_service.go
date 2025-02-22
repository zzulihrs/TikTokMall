package merchant

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/tiktokmall/backend/app/frontend/biz/service"
	"github.com/tiktokmall/backend/app/frontend/biz/utils"
	common "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/common"
	merchant "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/merchant"
)

// MerchantAddProduct .
// @router /merchant/product/add [POST]
func MerchantAddProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req merchant.MerchantAddProductReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, nil)
		// utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewMerchantAddProductService(ctx, c).Run(&req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, utils.H{
			"code":    500,
			"message": err.Error(),
		})
		// utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.JSON(consts.StatusOK, utils.WarpResponse(ctx, c, resp))
	// utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// MerchantDeleteProduct .
// @router /merchant/product/delete [POST]
func MerchantDeleteProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req merchant.MerchantDeleteProductReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, nil)
		// utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewMerchantDeleteProductService(ctx, c).Run(&req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, map[string]any{
			"code":    500,
			"message": err.Error(),
		})
		// utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.JSON(consts.StatusOK, utils.WarpResponse(ctx, c, resp))
	// utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// MerchantUpdateProduct .
// @router /merchant/product/update [POST]
func MerchantUpdateProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req merchant.MerchantUpdateProductReq
	// BUG: 无法修改类别
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, nil)
		// utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewMerchantUpdateProductService(ctx, c).Run(&req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, map[string]any{
			"code":    500,
			"message": err.Error(),
		})
		// utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.JSON(consts.StatusOK, utils.WarpResponse(ctx, c, resp))
	// utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// MerchantGetProductList .
// @router /merchant/product/list [POST]
func MerchantGetProductList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req merchant.MerchantGetProductListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, nil)
		// utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewMerchantGetProductListService(ctx, c).Run(&req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, map[string]any{
			"code":    500,
			"message": err.Error(),
		})
		// utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.JSON(consts.StatusOK, utils.WarpResponse(ctx, c, resp))
	// utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// MerchantGetProductDetail .
// @router /merchant/product/detail [POST]
func MerchantGetProductDetail(ctx context.Context, c *app.RequestContext) {
	var err error
	var req merchant.MerchantGetProductDetailReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, nil)
		// utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewMerchantGetProductDetailService(ctx, c).Run(&req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, map[string]any{
			"code":    500,
			"message": err.Error(),
		})
		// utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.JSON(consts.StatusOK, utils.WarpResponse(ctx, c, resp))
	// utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// MerchantPing .
// @router /merchant/ping [GET]
func MerchantPing(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, nil)
		return
	}

	resp, err := service.NewMerchantPingService(ctx, c).Run(&req)

	if err != nil {
		c.JSON(consts.StatusInternalServerError, map[string]any{
			"code":    500,
			"message": err.Error(),
		})
		return
	}
	c.JSON(consts.StatusOK, utils.WarpResponse(ctx, c, resp))
}

// MerchantRegister .
// @router /merchant/register [GET]
func MerchantRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	var req merchant.MerchantRegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, nil)
		return
	}

	resp, err := service.NewMerchantRegisterService(ctx, c).Run(&req)

	if err != nil {
		c.JSON(consts.StatusInternalServerError, map[string]any{
			"code":    500,
			"message": err.Error(),
		})
		return
	}
	c.JSON(consts.StatusOK, utils.WarpResponse(ctx, c, resp))
}

// MerchantAuth .
// @router /merchant/auth [GET]
func MerchantAuth(ctx context.Context, c *app.RequestContext) {
	var err error
	var req merchant.MerchantAuthReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, nil)
		// utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewMerchantAuthService(ctx, c).Run(&req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, map[string]any{
			"code":    500,
			"message": err.Error(),
		})
		// utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.JSON(consts.StatusOK, resp)
	// utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
