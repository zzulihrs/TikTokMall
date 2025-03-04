package category

import (
	"context"

	"github.com/tiktokmall/backend/app/frontend/biz/service"
	"github.com/tiktokmall/backend/app/frontend/biz/utils"
	category "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/category"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Category .
// @router /category/:category [GET]
func Category(ctx context.Context, c *app.RequestContext) {
	var err error
	var req category.CategoryReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, map[string]any{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	resp, err := service.NewCategoryService(ctx, c).Run(&req)
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
