package service

import (
	"context"
	"github.com/tiktokmall/backend/app/order/biz/dal/mysql"
	"github.com/tiktokmall/backend/app/order/biz/model"
	order "github.com/tiktokmall/backend/rpc_gen/kitex_gen/order"
)

type ChangeOrderStatusService struct {
	ctx context.Context
} // NewChangeOrderStatusService new ChangeOrderStatusService
func NewChangeOrderStatusService(ctx context.Context) *ChangeOrderStatusService {
	return &ChangeOrderStatusService{ctx: ctx}
}

// Run create note info
func (s *ChangeOrderStatusService) Run(req *order.ChangeOrderStatusReq) (resp *order.ChangeOrderStatusResp, err error) {
	err = model.ChangeOrderStatus(mysql.DB, s.ctx, req.OrderId, req.OrderStatus)
	return &order.ChangeOrderStatusResp{}, err
}
