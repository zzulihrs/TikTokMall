package payment

import (
	"context"
	"fmt"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/tiktokmall/backend/app/frontend/infra/rpc"
	frontendUtils "github.com/tiktokmall/backend/app/frontend/utils"
	rpcpayment "github.com/tiktokmall/backend/rpc_gen/kitex_gen/payment"
)

// PaymentMethod 定义支付方式
type PaymentMethod string

const (
	MethodAlipay PaymentMethod = "alipay"
	MethodWeChat PaymentMethod = "wechat"
	MethodCard   PaymentMethod = "card"
)

// PaymentRequest 请求结构体
type PaymentRequest struct {
	Method        PaymentMethod `json:"method" jsonschema:"description=Payment method to use"`
	UserID        uint32        `json:"user_id" jsonschema:"description=User ID"`
	Amount        float32       `json:"amount" jsonschema:"description=Payment amount"`
	CardInfo      *CardInfo     `json:"card_info,omitempty" jsonschema:"description=Card payment details"`
}

// CardInfo 银行卡信息
type CardInfo struct {
	CardNumber string `json:"card_number"`
	ExpiryDate string `json:"expiry_date"`
	CVV        string `json:"cvv"`
	HolderName string `json:"holder_name"`
}

// PaymentResponse 响应结构体
type PaymentResponse struct {
	Success     bool    `json:"success"`
	OrderID     uint32  `json:"order_id"`
	Amount      float32 `json:"amount"`
	Method      string  `json:"method"`
	PaymentURL  string  `json:"payment_url,omitempty"` // 支付宝/微信支付URL
	Transaction string  `json:"transaction,omitempty"` // 交易号
	Error       string  `json:"error,omitempty"`
}

// PaymentImpl 实现支付工具
type PaymentImpl struct {
	config *PaymentConfig
}

// PaymentConfig 配置结构体
type PaymentConfig struct{}

// NewPaymentTool 创建新的支付工具
func NewPaymentTool(ctx context.Context, config *PaymentConfig) (tool.BaseTool, error) {
	t := &PaymentImpl{config: config}
	return t.ToEinoTool()
}

// ToEinoTool 转换为 Eino 工具
func (p *PaymentImpl) ToEinoTool() (tool.BaseTool, error) {
	return utils.InferTool(
		"payment",
		// "Process payments using different payment methods",
		"Process payments only using alipay and get the url to open",
		p.Invoke,
	)
}

// Invoke 实际执行支付操作
func (p *PaymentImpl) Invoke(ctx context.Context, req *PaymentRequest) (*PaymentResponse, error) {
	req.UserID = frontendUtils.GetUserIdFromCtx(ctx)
	// 直接使用alipay，默认地址
	rpc.CheckoutClient.Checkout(ctx, &rpcpayment.CheckoutReq{
		UserId: req.UserID,
		Firstname: "test",
		Lastname: "test",
		Email: "root@example.com",
		PaymentMethod: req.Method,
		Address: &rpcpayment.Address{
			Country:       "China",
			ZipCode:       "200000",
			City:          "Shanghai",
			State:         "Shanghai",
			StreetAddress: "test",
		},}
	)
	
	return &PaymentResponse{
		Success:     true,
		PaymentURL: "http://tiktokmall.bhclient.cn/orders"，
	}
}