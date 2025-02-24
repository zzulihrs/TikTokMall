package productlist

import (
	"context"
	"fmt"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/tiktokmall/backend/app/frontend/infra/rpc"
	rpcproduct "github.com/tiktokmall/backend/rpc_gen/kitex_gen/product"
)

// Action 定义产品操作类型
type Action string
type Category string

const (
	ActionList Action = "list" // 列出所有产品
	ActionGet  Action = "get"  // 获取单个产品
)
const (
	Tshit   Category = "T-Shirt"
	Sticker Category = "Sticker"
)

// ProductRequest 请求结构体
type ProductRequest struct {
	Action   Action `json:"action" jsonschema:"description=Product action to perform"`
	ID       uint32 `json:"id,omitempty" jsonschema:"description=Product ID for get action"`
	Category string `json:"category,omitempty" jsonschema:"description=Catorgory for product list"`
}

// Product 产品信息
type Product struct {
	ID          uint32  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Category    string  `json:"category"`
}

// ProductResponse 响应结构体
type ProductResponse struct {
	Title    string    `json:"title"`
	Products []Product `json:"products"`
	Total    int       `json:"total"`
	Error    string    `json:"error,omitempty"`
}

// ProductImpl 实现产品查询工具
type ProductImpl struct {
	config *ProductConfig
}

// ProductConfig 配置结构体
type ProductConfig struct{}

// NewProduct 创建新的产品查询工具
func NewProduct(ctx context.Context, config *ProductConfig) (tool.BaseTool, error) {
	t := &ProductImpl{config: config}
	return t.ToEinoTool()
}

// ToEinoTool 转换为 Eino 工具
func (p *ProductImpl) ToEinoTool() (tool.BaseTool, error) {
	return utils.InferTool(
		"products",
		"Query and search products by Sticker or T-Shit catagory, respond with IDss",
		p.Invoke,
	)
}

// Invoke 实际执行产品查询操作
func (p *ProductImpl) Invoke(ctx context.Context, req *ProductRequest) (*ProductResponse, error) {
	res := &ProductResponse{
		Title: "Products",
	}

	switch req.Action {
	case ActionList:
		// 获取所有产品
		fmt.Printf("--------------------------------Product list request: %+v\n", req)
		productsResp, err := rpc.ProductClient.ListProducts(ctx, &rpcproduct.ListProductsReq{
			CategoryName: req.Category,
		})
		if err != nil {
			res.Error = fmt.Sprintf("Failed to list products: %v", err)
			return res, nil
		}
		return p.buildProductsResponse(productsResp.Products)

	case ActionGet:
		// 获取单个产品
		if req.ID == 0 {
			res.Error = "Product ID is required"
			return res, nil
		}
		productResp, err := rpc.ProductClient.GetProduct(ctx, &rpcproduct.GetProductReq{
			Id: req.ID,
		})
		if err != nil {
			res.Error = fmt.Sprintf("Failed to get product: %v", err)
			return res, nil
		}
		return p.buildSingleProductResponse(productResp.Product)
	}

	res.Error = "Invalid action"
	return res, nil
}

// buildProductsResponse 构建产品列表响应
func (p *ProductImpl) buildProductsResponse(products []*rpcproduct.Product) (*ProductResponse, error) {
	var productList []Product
	for _, prod := range products {
		productList = append(productList, Product{
			ID:          prod.Id,
			Name:        prod.Name,
			Description: prod.Description,
			Price:       prod.Price,
		})
	}

	fmt.Printf("--------------------------------Product list response: %+v\n", productList)
	return &ProductResponse{
		Title:    "Products",
		Products: productList,
		Total:    len(productList),
	}, nil
}

// buildSingleProductResponse 构建单个产品响应
func (p *ProductImpl) buildSingleProductResponse(prod *rpcproduct.Product) (*ProductResponse, error) {
	return &ProductResponse{
		Title: "Product Details",
		Products: []Product{{
			ID:          prod.Id,
			Name:        prod.Name,
			Description: prod.Description,
			Price:       prod.Price,
		}},
		Total: 1,
	}, nil
}
