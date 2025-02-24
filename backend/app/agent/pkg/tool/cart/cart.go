package cart

import (
	"context"
	"fmt"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/tiktokmall/backend/app/frontend/infra/rpc"
	rpccart "github.com/tiktokmall/backend/rpc_gen/kitex_gen/cart"
)

// Action 定义购物车操作类型
type Action string

const (
	ActionAdd    Action = "add"    // 添加商品到购物车
	ActionGet    Action = "get"    // 获取购物车信息
	ActionEmpty  Action = "empty"  // 清空购物车
	ActionUpdate Action = "update" // 更新商品数量
)

// CartRequest 请求结构体
type CartRequest struct {
	Action   Action `json:"action" jsonschema:"description=Cart action to perform"`
	UserID   uint32 `json:"user_id" jsonschema:"description=User ID"`
	ItemID   uint32 `json:"item_id,omitempty" jsonschema:"description=Item ID for add/update actions"`
	Quantity int32  `json:"quantity,omitempty" jsonschema:"description=Quantity for add/update actions"`
}

// CartItem 购物车项目
type CartItem struct {
	ProductID   uint32  `json:"product_id"`
	ProductName string  `json:"product_name"`
	Picture     string  `json:"picture"`
	Price       float32 `json:"price"`
	Quantity    int32   `json:"quantity"`
}

// CartResponse 响应结构体
type CartResponse struct {
	Title   string     `json:"title"`
	Items   []CartItem `json:"items"`
	Total   float32    `json:"total"`
	CartNum int        `json:"cart_num"`
	Error   string     `json:"error,omitempty"`
}

// CartImpl 实现购物车工具
type CartImpl struct {
	config *CartConfig
}

// CartConfig 配置结构体
type CartConfig struct{}

// NewCart 创建新的购物车工具
func NewCart(ctx context.Context, config *CartConfig) (tool.BaseTool, error) {
	t := &CartImpl{config: config}
	return t.ToEinoTool()
}

// ToEinoTool 转换为 Eino 工具
func (c *CartImpl) ToEinoTool() (tool.BaseTool, error) {
	return utils.InferTool(
		"cart",
		"Manage shopping cart operations",
		c.Invoke,
	)
}

// Invoke 实际执行购物车操作
func (c *CartImpl) Invoke(ctx context.Context, req *CartRequest) (*CartResponse, error) {
	res := &CartResponse{
		Title: "Cart",
	}

	switch req.Action {
	case ActionAdd:
		_, err := rpc.CartClient.AddToCart(ctx, &rpccart.AddToCartReq{
			UserId:    req.UserID,
			ProductId: req.ItemID,
			Quantity:  req.Quantity,
		})
		if err != nil {
			res.Error = fmt.Sprintf("Failed to add item: %v", err)
			return res, nil
		}

	case ActionGet:
		cartResp, err := rpc.CartClient.GetCart(ctx, &rpccart.GetCartReq{
			UserId: req.UserID,
		})
		if err != nil {
			res.Error = fmt.Sprintf("Failed to get cart: %v", err)
			return res, nil
		}
		return c.buildCartResponse(ctx, cartResp)

	case ActionEmpty:
		_, err := rpc.CartClient.EmptyCart(ctx, &rpccart.EmptyCartReq{
			UserId: req.UserID,
		})
		if err != nil {
			res.Error = fmt.Sprintf("Failed to empty cart: %v", err)
			return res, nil
		}

	case ActionUpdate:
		_, err := rpc.CartClient.UpdateCart(ctx, &rpccart.UpdateCartReq{
			UserId:    req.UserID,
			ProductId: req.ItemID,
			Quantity:  req.Quantity,
		})
		if err != nil {
			res.Error = fmt.Sprintf("Failed to update quantity: %v", err)
			return res, nil
		}
	}

	// 获取最新的购物车状态
	return c.getUpdatedCart(ctx, req.UserID)
}

// buildCartResponse 构建购物车响应
func (c *CartImpl) buildCartResponse(ctx context.Context, cartResp *rpccart.GetCartResp) (*CartResponse, error) {
	var items []CartItem
	var total float32

	for _, item := range cartResp.Cart.Items {
		productResp, err := rpc.ProductClient.GetProduct(ctx, &rpcproduct.GetProductReq{
			Id: item.ProductId,
		})
		if err != nil {
			continue
		}

		items = append(items, CartItem{
			ProductID:   item.ProductId,
			ProductName: productResp.Product.Name,
			Picture:     productResp.Product.Picture,
			Price:       productResp.Product.Price,
			Quantity:    item.Quantity,
		})
		total += productResp.Product.Price * float32(item.Quantity)
	}

	return &CartResponse{
		Title:   "Cart",
		Items:   items,
		Total:   total,
		CartNum: len(items),
	}, nil
}

// getUpdatedCart 获取更新后的购物车状态
func (c *CartImpl) getUpdatedCart(ctx context.Context, userID uint32) (*CartResponse, error) {
	cartResp, err := rpc.CartClient.GetCart(ctx, &rpccart.GetCartReq{
		UserId: userID,
	})
	if err != nil {
		return &CartResponse{
			Title: "Cart",
			Error: fmt.Sprintf("Failed to get updated cart: %v", err),
		}, nil
	}
	return c.buildCartResponse(ctx, cartResp)
}
