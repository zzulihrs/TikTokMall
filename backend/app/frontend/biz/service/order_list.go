package service

import (
	"context"
	"time"

	common "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/common"
	"github.com/tiktokmall/backend/app/frontend/infra/rpc"
	"github.com/tiktokmall/backend/app/frontend/types"
	frontendUtils "github.com/tiktokmall/backend/app/frontend/utils"
	rpcorder "github.com/tiktokmall/backend/rpc_gen/kitex_gen/order"
	rpcproduct "github.com/tiktokmall/backend/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/app"
)

type OrderListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderListService(Context context.Context, RequestContext *app.RequestContext) *OrderListService {
	return &OrderListService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderListService) Run(req *common.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	userId := frontendUtils.GetUserIdFromCtx(h.Context)
	orderResp, err := rpc.OrderClient.ListOder(h.Context, &rpcorder.ListOrderReq{UserId: userId})
	if err != nil {
		return nil, err
	}
	var list []*types.Order
	for _, v := range orderResp.Orders {
		var (
			total float32
			items []types.OrderItem
		)
		for _, item := range v.Items {
			total += item.Cost
			productResp, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{Id: item.Item.ProductId})
			if err != nil {
				return nil, err
			}
			if productResp == nil || productResp.Product == nil {
				continue
			}
			p := productResp.Product
			i := item.Item
			items = append(items, types.OrderItem{
				ProductName: p.Name,
				Picture:     p.Picture,
				Cost:        item.Cost,
				Qty:         int32(i.Quantity),
			})
		}

		created := time.Unix(int64(v.CreatedAt), 0)
		list = append(list, &types.Order{
			OrderId:     v.OrderId,
			CreatedDate: created.Format("2006-01-02 15:04:05"),
			Cost:        total,
			Items:       items,
		})
	}
	return map[string]any{
		"title":  "order",
		"orders": list,
	}, nil
}
