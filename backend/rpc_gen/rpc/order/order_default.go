package order

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	order "github.com/tiktokmall/backend/rpc_gen/kitex_gen/order"
)

func PlaceOrder(ctx context.Context, req *order.PlaceOrderReq, callOptions ...callopt.Option) (resp *order.PlaceOrderResp, err error) {
	resp, err = defaultClient.PlaceOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "PlaceOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ListOder(ctx context.Context, req *order.ListOrderReq, callOptions ...callopt.Option) (resp *order.ListOrderResp, err error) {
	resp, err = defaultClient.ListOder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ListOder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ChangeOrderStatus(ctx context.Context, req *order.ChangeOrderStatusReq, callOptions ...callopt.Option) (resp *order.ChangeOrderStatusResp, err error) {
	resp, err = defaultClient.ChangeOrderStatus(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ChangeOrderStatus call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
