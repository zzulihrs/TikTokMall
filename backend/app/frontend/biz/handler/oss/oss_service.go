package oss

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/tiktokmall/backend/app/frontend/biz/service"
	"github.com/tiktokmall/backend/app/frontend/biz/utils"
	oss "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/oss"
)

// UploadImage .
// @router /uploadImage [POST]
func UploadImage(ctx context.Context, c *app.RequestContext) {

	var req oss.UploadFileRequest
	resp, err := service.NewUploadImageService(ctx, c).Run(&req)

	if err != nil {
		c.JSON(consts.StatusInternalServerError, map[string]any{
			"code":    500,
			"message": err.Error(),
		})
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UploadVideo .
// @router /uploadVideo [POST]
func UploadVideo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req oss.UploadFileRequest

	resp, err := service.NewUploadVideoService(ctx, c).Run(&req)

	if err != nil {
		c.JSON(consts.StatusInternalServerError, map[string]any{
			"code":    500,
			"message": err.Error(),
		})
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
