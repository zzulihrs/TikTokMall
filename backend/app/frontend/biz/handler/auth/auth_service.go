package auth

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/tiktokmall/backend/app/frontend/biz/service"
	"github.com/tiktokmall/backend/app/frontend/biz/utils"
	auth "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/auth"
	common "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/common"
	mw "github.com/tiktokmall/backend/app/frontend/middleware"
)

// Login .
// @router /auth/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	// var err error
	// var req auth.LoginReq
	// err = c.BindAndValidate(&req)
	// if err != nil {
	// 	utils.SendErrResponse(ctx, c, consts.StatusOK, err)
	// 	return
	// }

	// redirect, err := service.NewLoginService(ctx, c).Run(&req)
	// if err != nil {
	// 	c.JSON(consts.StatusAccepted, "账号不存在/账号密码不匹配")
	// 	//utils.SendErrResponse(ctx, c, consts.StatusOK, err)
	// 	return
	// }

	// c.JSON(consts.StatusOK, redirect)
	mw.JwtMiddleware.LoginHandler(ctx, c)

	//c.Redirect(consts.StatusOK, []byte(redirect))

	// utils.SendSuccessResponse(ctx, c, consts.StatusOK, "done!")
}

// Register .
// @router /auth/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.RegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	user_id, err := service.NewRegisterService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	token, _, err := mw.JwtMiddleware.TokenGenerator(user_id)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.JSON(consts.StatusOK, &utils.H{
		"redirect": "login",
		"token":    token,
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
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	_, err = service.NewLogoutService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.JSON(consts.StatusOK, "/")

	//c.Redirect(consts.StatusOK, []byte("/"))
	// utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
