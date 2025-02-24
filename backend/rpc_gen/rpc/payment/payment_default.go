package payment

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	payment "github.com/tiktokmall/backend/rpc_gen/kitex_gen/payment"
)

func Charge(ctx context.Context, req *payment.ChargeReq, callOptions ...callopt.Option) (resp *payment.ChargeResp, err error) {
	resp, err = defaultClient.Charge(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Charge call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Alipay(ctx context.Context, req *payment.AlipayReq, callOptions ...callopt.Option) (resp *payment.AlipayResp, err error) {
	resp, err = defaultClient.Alipay(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Alipay call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
