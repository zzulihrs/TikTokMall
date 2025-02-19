// Code generated by hertz generator.

package merchant

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/tiktokmall/backend/app/frontend/biz/utils"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _merchantMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _merchantauthMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _merchantpingMw() []app.HandlerFunc {
	// your code...
	return nil
}

// 保护 merchant/product 下的所有路由
func _productMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{func(ctx context.Context, c *app.RequestContext) {
		bs := c.GetHeader("Authorization")
		if !utils.CheckToken(bs) {
			c.JSON(consts.StatusBadRequest, utils.H{
				"code":    400,
				"message": "token invalid",
			})
			c.Abort()
			return
		}
	}}
}

func _merchantaddproductMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _merchantdeleteproductMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _merchantgetproductdetailMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _merchantgetproductlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _merchantupdateproductMw() []app.HandlerFunc {
	// your code...
	return nil
}
