package payment

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/tiktokmall/backend/app/frontend/biz/service"
	"github.com/tiktokmall/backend/app/frontend/biz/utils"
	common "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/common"
	payment "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/payment"
)

// Alipay .
// @router /alipay [GET]
func Alipay(ctx context.Context, c *app.RequestContext) {
	var err error
	var req payment.AlipayReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewAlipayService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.JSON(consts.StatusOK, utils.WarpResponse(ctx, c, resp))

}

// Payresult .
// @router /payresult [POST]
func Payresult(ctx context.Context, c *app.RequestContext) {
	var err error
	var req payment.PayresultReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewPayresultService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	//utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
	c.JSON(consts.StatusOK, utils.WarpResponse(ctx, c, resp))
}

// Paysuccess .
// @router /paysuccess [GET]
func Paysuccess(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewPaysuccessService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	//utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
	// 需要提供一个显示支付成功的页面，跳转到该页面
	c.JSON(consts.StatusOK, utils.WarpResponse(ctx, c, resp))
}
