package service

import (
	"context"
	"fmt"
	"github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/payment"
	"github.com/tiktokmall/backend/app/frontend/infra/rpc"
	rpcpayment "github.com/tiktokmall/backend/rpc_gen/kitex_gen/payment"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
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
	alipayResp, err := rpc.PaymentClient.Alipay(h.Context, &rpcpayment.AlipayReq{
		TransactionId: strconv.Itoa(int(req.GetTransactionId())),
		TotalAmount:   float32(req.GetTotalAmount()),
	})

	// 返回的PayUrl是一个网址，需要前端打开该网址
	fmt.Println(alipayResp.GetPayUrl())
	return map[string]any{
		"PayUrl": alipayResp.GetPayUrl(),
	}, err
}
