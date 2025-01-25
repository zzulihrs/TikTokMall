package product

import (
	"context"
	"log"

	"github.com/tiktokmall/backend/app/frontend/biz/service"
	"github.com/tiktokmall/backend/app/frontend/biz/utils"
	product "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/product"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// GetProduct .
// @router /product [GET]
func GetProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.ProductReq
	err = c.BindAndValidate(&req)
	log.Printf("product, args: %#v\n", req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	log.Printf("GetProduct\n")

	resp, err := service.NewGetProductService(ctx, c).Run(&req)
	log.Printf("GetProduct, Resp: %v, err: %v\n", resp, err)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// c.HTML(consts.StatusOK, "product", utils.WarpResponse(ctx, c, resp))
	c.JSON(consts.StatusOK, utils.WarpResponse(ctx, c, resp))
}

// SearchProducts .
// @router /search [GET]
func SearchProducts(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.SearchProductsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewSearchProductsService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// c.HTML(consts.StatusOK, "search", utils.WarpResponse(ctx, c, resp))
	c.JSON(consts.StatusOK, utils.WarpResponse(ctx, c, resp))
}
