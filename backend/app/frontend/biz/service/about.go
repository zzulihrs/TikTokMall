package service

import (
	"context"

	common "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/common"

	"github.com/cloudwego/hertz/pkg/app"
)

type AboutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAboutService(Context context.Context, RequestContext *app.RequestContext) *AboutService {
	return &AboutService{RequestContext: RequestContext, Context: Context}
}

func (h *AboutService) Run(req *common.Empty) (resp map[string]any, err error) {
	return map[string]any{
		"title": "About",
	}, nil
}
