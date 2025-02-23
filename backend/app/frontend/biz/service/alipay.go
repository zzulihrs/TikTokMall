package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/payment"
	"github.com/tiktokmall/backend/app/frontend/infra/rpc"
	rpcpayment "github.com/tiktokmall/backend/rpc_gen/kitex_gen/payment"
	"log"
)

type AlipayService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAlipayService(Context context.Context, RequestContext *app.RequestContext) *AlipayService {
	return &AlipayService{RequestContext: RequestContext, Context: Context}
}

func (h *AlipayService) Run(req *payment.AlipayReq) (resp map[string]any, err error) {
	// 远程调用alipay功能
	log.Println("out_trade_no", req.TransactionId)
	log.Println("total_amount", req.TotalAmount)

	alipayResp, err := rpc.PaymentClient.Alipay(h.Context, &rpcpayment.AlipayReq{
		TransactionId: req.GetTransactionId(),
		TotalAmount:   req.GetTotalAmount(),
	})

	// 返回的PayUrl是一个网址，需要前端打开该网址
	fmt.Println(alipayResp.GetPayUrl())
	return map[string]any{
		"PayUrl": alipayResp.GetPayUrl(),
	}, err
}
