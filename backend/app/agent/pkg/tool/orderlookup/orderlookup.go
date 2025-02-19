package orderlookup

import (
	"context"
	"fmt"

	"time"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/tiktokmall/backend/app/frontend/infra/rpc"
	frontendUtils "github.com/tiktokmall/backend/app/frontend/utils"
	rpcorder "github.com/tiktokmall/backend/rpc_gen/kitex_gen/order"
	rpcproduct "github.com/tiktokmall/backend/rpc_gen/kitex_gen/product"
)

// OrderLookupImpl 实现订单查询工具
type OrderLookupImpl struct {
	config *OrderLookupConfig
}

// OrderLookupConfig 配置结构体
type OrderLookupConfig struct {
	// TODO：添加配置项
}

// OrderLookupRequest 请求结构体
type OrderLookupRequest struct {
	UserID uint32 `json:"user_id" jsonschema:"description=User ID to look up orders for"`
}

// OrderItem 订单项
type OrderItem struct {
	ProductName string  `json:"product_name"`
	Picture     string  `json:"picture"`
	Cost        float32 `json:"cost"`
	Qty         int32   `json:"qty"`
}

// Order 订单信息
type Order struct {
	OrderId     string      `json:"order_id"`
	CreatedDate string      `json:"created_date"`
	Cost        float32     `json:"cost"`
	Items       []OrderItem `json:"items"`
}

// OrderLookupResponse 响应结构体
type OrderLookupResponse struct {
	Title  string   `json:"title"`
	Orders []*Order `json:"orders"`
	Error  string   `json:"error,omitempty"`
}

// NewOrderLookup 创建新的订单查询工具
func NewOrderLookup(ctx context.Context, config *OrderLookupConfig) (tool.BaseTool, error) {
	t := &OrderLookupImpl{config: config}
	return t.ToEinoTool()
}

// ToEinoTool 转换为 Eino 工具
func (o *OrderLookupImpl) ToEinoTool() (tool.BaseTool, error) {
	return utils.InferTool(
		"orderlookup",
		"Look up orders for a user",
		o.Invoke,
	)
}

// Invoke 实际执行订单查询
func (o *OrderLookupImpl) Invoke(ctx context.Context, req *OrderLookupRequest) (*OrderLookupResponse, error) {
	res := &OrderLookupResponse{
		Title: "order",
	}

	req.UserID = frontendUtils.GetUserIdFromCtx(ctx)
	// TODO：error：id always 0
	req.UserID = 15
	if req.UserID == 0 {
		res.Error = "UserID cannot be empty"
		return res, nil
	}

	// 调用 RPC 客户端查询订单
	orderResp, err := rpc.OrderClient.ListOder(ctx, &rpcorder.ListOrderReq{
		UserId: req.UserID,
	})
	if err != nil {
		res.Error = fmt.Sprintf("Failed to fetch orders: %v", err)
		return res, nil
	}

	// 处理订单数据
	var orders []*Order
	for _, v := range orderResp.Orders {
		var (
			total float32
			items []OrderItem
		)

		// 处理订单项
		for _, item := range v.Items {
			total += item.Cost
			// 获取商品详情
			productResp, err := rpc.ProductClient.GetProduct(ctx, &rpcproduct.GetProductReq{
				Id: item.Item.ProductId,
			})
			if err != nil {
				continue
			}
			if productResp == nil || productResp.Product == nil {
				continue
			}

			p := productResp.Product
			i := item.Item
			items = append(items, OrderItem{
				ProductName: p.Name,
				Picture:     p.Picture,
				Cost:        item.Cost,
				Qty:         int32(i.Quantity),
			})
		}

		// 格式化创建时间
		created := time.Unix(int64(v.CreatedAt), 0)

		// 构建订单数据
		orders = append(orders, &Order{
			OrderId:     v.OrderId,
			CreatedDate: created.Format("2006-01-02 15:04:05"),
			Cost:        total,
			Items:       items, // 太长了一般用不上
		})
	}

	res.Orders = orders
	return res, nil
}
