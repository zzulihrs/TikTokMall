package user

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/tiktokmall/backend/app/frontend/biz/service"
	"github.com/tiktokmall/backend/app/frontend/biz/utils"
	user "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/user"
)

// Update .
// @router /user/update [POST]
func Update(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UpdateUserReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &user.UpdateUserResp{}
	_, err = service.NewUpdateService(ctx, c).Run(&req)
	if err != nil {
		c.JSON(consts.StatusAccepted, &user.UpdateUserResp{
			StatusCode: consts.StatusInternalServerError,
			Msg:        err.Error(),
		})
		//utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	resp.StatusCode = consts.StatusOK
	resp.Msg = "update successfully"
	c.JSON(consts.StatusOK, resp)
}

// DeleteUser .
// @router /user/delete/:user_id [POST]
func DeleteUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.DeleteUserReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &user.DeleteUserResp{}
	resp, err = service.NewDeleteUserService(ctx, c).Run(&req)
	if err != nil {
		c.JSON(consts.StatusAccepted, &user.DeleteUserResp{
			StatusCode: consts.StatusInternalServerError,
			Msg:        err.Error(),
		})
		return
	}
	c.JSON(consts.StatusOK, resp)
}

// QueryUser .
// @router /user/query/:user_id [GET]
func QueryUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.QueryUserReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	_, err = service.NewQueryUserService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	//utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
	c.JSON(consts.StatusOK, "/query")
}
