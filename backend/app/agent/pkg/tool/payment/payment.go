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
	TransactionID string        `json:"order_id" jsonschema:"description=Transaction ID to pay for"`
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
	// TODO: 灵活处理transactionid
	req.TransactionID = "c3332f08-f2aa-11ef-b7d1-525400e93088"
	// 根据支付方式进行不同处理
	switch req.Method {
	case MethodAlipay:
		return p.handleAlipayPayment(ctx, req)
	// case MethodWeChat:
	// 	return p.handleWeChatPayment(ctx, req)
	// case MethodCard:
	// 	return p.handleCardPayment(ctx, req)
	default:
		return &PaymentResponse{
			Success: false,
			Error:   "unsupported payment method",
		}, nil
	}
}

// 处理支付宝支付
func (p *PaymentImpl) handleAlipayPayment(ctx context.Context, req *PaymentRequest) (*PaymentResponse, error) {
	payReq := &rpcpayment.AlipayReq{
		TransactionId: req.TransactionID,
		TotalAmount:   fmt.Sprintf("%.2f", req.Amount),
	}

	return p.processPayment(ctx, payReq)
}

// // 处理微信支付
// func (p *PaymentImpl) handleWeChatPayment(ctx context.Context, req *PaymentRequest) (*PaymentResponse, error) {
// 	// 等待后续。。。
// 	payReq := &rpcpayment.ChargeReq{}

// 	return p.processPayment(ctx, payReq)
// }

// // 处理银行卡支付
// func (p *PaymentImpl) handleCardPayment(ctx context.Context, req *PaymentRequest) (*PaymentResponse, error) {
// 	// 等待后续。。。
// 	payReq := rpcpayment.ChargeReq{}

// 	return p.processPayment(ctx, payReq)
// }

// 统一处理支付流程
func (p *PaymentImpl) processPayment(ctx context.Context, payReq *rpcpayment.AlipayReq) (*PaymentResponse, error) {
	payResp, err := rpc.PaymentClient.Alipay(ctx, payReq)
	if err != nil {
		return &PaymentResponse{
			Success: false,
			Error:   fmt.Sprintf("Failed to create payment: %v", err),
		}, nil
	}

	resp := &PaymentResponse{}

	// 如果是在线支付方式，设置支付URL
	// if payReq.Method == string(MethodAlipay) || payReq.Method == string(MethodWeChat) {
	// 	resp.PaymentURL = payResp.PaymentUrl
	// }
	resp.PaymentURL = payResp.GetPayUrl()
	return resp, nil
}
