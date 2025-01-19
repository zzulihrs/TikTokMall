package service

import (
	"context"
	"strconv"

	common "bytetech/course/app/frontend/hertz_gen/frontend/common"
	"bytetech/course/app/frontend/infra/rpc"
	frontendUtils "bytetech/course/app/frontend/utils"
	rpccart "bytetech/course/rpc_gen/kitex_gen/cart"
	rpcproduct "bytetech/course/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/app"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(req *common.Empty) (resp map[string]any, err error) {
	carts, err := rpc.CartClient.GetCart(h.Context, &rpccart.GetCartReq{
		UserId: frontendUtils.GetUserIdFromCtx(h.Context),
	})
	if err != nil {
		return nil, err
	}
	// total 表示总价
	var items []map[string]string
	var total float32
	for _, item := range carts.Cart.Items {
		productResp, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{Id: item.GetProductId()})
		if err != nil {
			continue
		}
		if productResp.Product == nil {
			continue
		}
		p := productResp.Product
		items = append(items, map[string]string{"Name": p.Name, "Description": p.Description, "Picture": p.Picture, "Price": strconv.FormatFloat(float64(p.Price), 'f', 2, 64), "Qty": strconv.Itoa(int(item.Quantity))})
		total += float32(item.Quantity) * p.Price
	}

	return map[string]any{
		"title": "Cart",
		"items": items,
		"total": strconv.FormatFloat(float64(total), 'f', 2, 64),
	}, nil
}
