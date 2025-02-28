package auth

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/tiktokmall/backend/app/frontend/biz/service"
	"github.com/tiktokmall/backend/app/frontend/biz/utils"
	auth "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/auth"
	common "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/common"
)

// Login .
// @router /auth/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.LoginReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewLoginService(ctx, c).Run(&req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, map[string]any{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// Register .
// @router /auth/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.RegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]any{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	_, err = service.NewRegisterService(ctx, c).Run(&req)

	if err != nil {
		c.JSON(consts.StatusInternalServerError, map[string]any{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	// token, _, err := mw.JwtMiddleware.TokenGenerator(user_id)
	// if err != nil {
	// 	utils.SendErrResponse(ctx, c, consts.StatusOK, err)
	// 	return
	// }

	c.JSON(consts.StatusOK, map[string]any{
		"code":    200,
		"message": "OK",
		"data": map[string]any{
			"email": req.Email,
		},
	})

	//c.Redirect(consts.StatusOK, []byte("/"))
	// utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// Logout .
// @router /auth/logout [POST]
func Logout(ctx context.Context, c *app.RequestContext) {
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

	_, err = service.NewLogoutService(ctx, c).Run(&req)

	if err != nil {
		c.JSON(consts.StatusInternalServerError, map[string]any{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	c.JSON(consts.StatusOK, map[string]any{
		"code":    200,
		"message": "OK",
	})

	//c.Redirect(consts.StatusOK, []byte("/"))
	// utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
