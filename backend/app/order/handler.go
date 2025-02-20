package main

import (
	"context"

	"github.com/tiktokmall/backend/app/order/biz/service"
	order "github.com/tiktokmall/backend/rpc_gen/kitex_gen/order"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// PlaceOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) PlaceOrder(ctx context.Context, req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	resp, err = service.NewPlaceOrderService(ctx).Run(req)

	return resp, err
}

// ListOder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) ListOder(ctx context.Context, req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	resp, err = service.NewListOderService(ctx).Run(req)

	return resp, err
}

// ChangeOrderStatus implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) ChangeOrderStatus(ctx context.Context, req *order.ChangeOrderStatusReq) (resp *order.ChangeOrderStatusResp, err error) {
	resp, err = service.NewChangeOrderStatusService(ctx).Run(req)

	return resp, err
}
