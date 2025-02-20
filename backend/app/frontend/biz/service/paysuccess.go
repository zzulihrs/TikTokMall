package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	common "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/common"
)

type PaysuccessService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewPaysuccessService(Context context.Context, RequestContext *app.RequestContext) *PaysuccessService {
	return &PaysuccessService{RequestContext: RequestContext, Context: Context}
}

func (h *PaysuccessService) Run(req *common.Empty) (resp map[string]any, err error) {
	// 直接跳转到支付成功界面
	return map[string]any{}, nil
}
