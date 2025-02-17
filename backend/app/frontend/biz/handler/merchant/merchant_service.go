package merchant

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/tiktokmall/backend/app/frontend/biz/service"
	"github.com/tiktokmall/backend/app/frontend/biz/utils"
	merchant "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/merchant"
)

// MerchantAuth .
// @router /marchant/auth [POST]
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
		c.JSON(consts.StatusInternalServerError, utils.WarpResponse(ctx, c, resp))
		// utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.JSON(consts.StatusOK, utils.WarpResponse(ctx, c, resp))
	// utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// MerchantAddProduct .
// @router /marchant/product/add [POST]
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
		c.JSON(consts.StatusInternalServerError, utils.WarpResponse(ctx, c, resp))
		// utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.JSON(consts.StatusOK, utils.WarpResponse(ctx, c, resp))
	// utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// MerchantDeleteProduct .
// @router /marchant/product/delete [POST]
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
		c.JSON(consts.StatusInternalServerError, utils.WarpResponse(ctx, c, resp))
		// utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.JSON(consts.StatusOK, utils.WarpResponse(ctx, c, resp))
	// utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// MerchantUpdateProduct .
// @router /marchant/product/update [POST]
func MerchantUpdateProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req merchant.MerchantUpdateProductReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, nil)
		// utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewMerchantUpdateProductService(ctx, c).Run(&req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, utils.WarpResponse(ctx, c, resp))
		// utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.JSON(consts.StatusOK, utils.WarpResponse(ctx, c, resp))
	// utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// MerchantGetProductList .
// @router /marchant/product/list [POST]
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
		c.JSON(consts.StatusInternalServerError, utils.WarpResponse(ctx, c, resp))
		// utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.JSON(consts.StatusOK, utils.WarpResponse(ctx, c, resp))
	// utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// MerchantGetProductDetail .
// @router /marchant/product/detail [POST]
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
		c.JSON(consts.StatusInternalServerError, utils.WarpResponse(ctx, c, resp))
		// utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.JSON(consts.StatusOK, utils.WarpResponse(ctx, c, resp))
	// utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
