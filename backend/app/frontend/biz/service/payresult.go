package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	payment "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/payment"
)

type PayresultService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewPayresultService(Context context.Context, RequestContext *app.RequestContext) *PayresultService {
	return &PayresultService{RequestContext: RequestContext, Context: Context}
}

func (h *PayresultService) Run(req *payment.PayresultReq) (resp map[string]any, err error) {
	// 需要修改订单状态---等待提供修改订单状态的接口

	return map[string]any{
		"message": "Success",
	}, nil
}
