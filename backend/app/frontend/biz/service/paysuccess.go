package service

import (
	"context"
	"fmt"
	"github.com/tiktokmall/backend/app/frontend/infra/rpc"
	"github.com/tiktokmall/backend/rpc_gen/kitex_gen/order"
	"log"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

type PaysuccessService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewPaysuccessService(Context context.Context, RequestContext *app.RequestContext) *PaysuccessService {
	return &PaysuccessService{RequestContext: RequestContext, Context: Context}
}

func (h *PaysuccessService) Run(req *http.Request) (resp map[string]any, err error) {
	// 需要修改订单状态---等待提供修改订单状态的接口
	// 解析异步通知的参数

	out_trade_no := h.RequestContext.Query("out_trade_no")
	log.Println("out_trade_no: ", out_trade_no)

	_, err = rpc.OrderClient.ChangeOrderStatus(h.Context, &order.ChangeOrderStatusReq{
		OrderId:     out_trade_no,
		OrderStatus: 1, // 订单状态 0 - 创建订单待支付 1 - 支付成功 2 - 支付失败 3 - 取消订单
	})

	if err != nil {
		return nil, fmt.Errorf("change order status failed: %v", err)
	}
	// 直接跳转到支付成功界面
	return map[string]any{}, nil
}
