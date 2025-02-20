package service

import (
	"context"
	"fmt"

	"github.com/tiktokmall/backend/app/checkout/infra/mq"
	"github.com/tiktokmall/backend/app/checkout/infra/rpc"
	"github.com/tiktokmall/backend/rpc_gen/kitex_gen/cart"
	checkout "github.com/tiktokmall/backend/rpc_gen/kitex_gen/checkout"
	"github.com/tiktokmall/backend/rpc_gen/kitex_gen/email"
	"github.com/tiktokmall/backend/rpc_gen/kitex_gen/order"
	"github.com/tiktokmall/backend/rpc_gen/kitex_gen/payment"
	"github.com/tiktokmall/backend/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/protobuf/proto"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
/*
	Run
1. get cart
2. calculate cart
3. create order
4. empty cart
5. pay
6. change order result
7. finish
*/
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	cartResult, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: req.GetUserId()})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5005001, err.Error())
	}
	if cartResult == nil || cartResult.Cart == nil || len(cartResult.Cart.Items) == 0 {
		return nil, kerrors.NewGRPCBizStatusError(5004001, "cart is empty")
	}
	var (
		total float32
		oi    []*order.OrderItem
	)
	for _, cartItem := range cartResult.Cart.Items {
		productResp, resultErr := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: cartItem.ProductId})
		if resultErr != nil {
			return nil, resultErr
		}
		if productResp == nil || productResp.Product == nil {
			continue
		}
		cost := productResp.Product.Price
		total += cost * float32(cartItem.Quantity)
		oi = append(oi, &order.OrderItem{
			Item: &cart.CartItem{
				ProductId: cartItem.ProductId,
				Quantity:  cartItem.Quantity,
			},
			Cost: cost,
		})
	}

	var orderId string
	// u, _ := uuid.NewRandom()
	// orderId = u.String()

	orderResp, err := rpc.OrderClient.PlaceOrder(s.ctx, &order.PlaceOrderReq{
		UserId: req.GetUserId(),
		Email:  req.GetEmail(),
		Address: &order.Address{
			StreetAddress: req.GetAddress().GetStreetAddress(),
			City:          req.GetAddress().GetCity(),
			State:         req.GetAddress().GetState(),
			Country:       req.GetAddress().GetCountry(),
			ZipCode:       req.GetAddress().GetZipCode(),
		},
		Items: oi,
	})

	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5004002, err.Error())
	}

	if orderResp == nil || orderResp.Order == nil {
		return nil, kerrors.NewGRPCBizStatusError(5004003, "orderId is nil")
	}

	orderId = orderResp.Order.OrderId

	payReq := payment.ChargeReq{
		Amount:  total,
		UserId:  req.GetUserId(),
		OrderId: orderId,
		CreditCard: &payment.CreditCardInfo{
			CreditCartNumber:          req.GetCreditCard().GetCreditCartNumber(),
			CreditCardCvv:             req.GetCreditCard().GetCreditCardCvv(),
			CreditCardExpirationYear:  req.GetCreditCard().GetCreditCardExpirationYear(),
			CreditCardExpirationMonth: req.GetCreditCard().GetCreditCardExpirationMonth(),
		},
	}

	paymentResult := &payment.ChargeResp{}
	if req.GetPaymentMethod() == "card" { // 银行卡支付
		paymentResult, err = rpc.PaymentClient.Charge(s.ctx, &payReq)
		if err != nil {
			return nil, fmt.Errorf("charge failed: %v", err)
		}
		// 修改订单状态为支付成功
		rpc.OrderClient.ChangeOrderStatus(s.ctx, &order.ChangeOrderStatusReq{
			OrderId:     orderId,
			OrderStatus: 1,
			UserId:      req.GetUserId(),
		})
	} else if req.GetPaymentMethod() == "card" { // 支付宝支付
		// TODO
	}

	_, err = rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.GetUserId()})
	if err != nil {
		return nil, fmt.Errorf("empty cart failed: %v", err)
	}

	// 发送邮件通知
	data, _ := proto.Marshal(&email.EmailReq{
		From:        "from@example.com",
		To:          req.Email,
		ContextType: "text/plain",
		Subject:     "You have juse created an order in the CloudWeGo Shop",
		Content:     "You have juse created an order in the CloudWeGo Shop",
	})
	msg := &nats.Msg{Subject: "email", Data: data, Header: make(nats.Header)}
	// 将上下文注入到 header 中
	otel.GetTextMapPropagator().Inject(s.ctx, propagation.HeaderCarrier(msg.Header))
	_ = mq.Nc.PublishMsg(msg)

	klog.Info(paymentResult)

	resp = &checkout.CheckoutResp{
		OrderId:       orderId,
		TransactionId: paymentResult.TransactionId,
	}
	return resp, nil
}
