package service

import (
	"context"

	"github.com/tiktokmall/backend/app/order/biz/dal/mysql"
	"github.com/tiktokmall/backend/app/order/biz/model"
	"github.com/tiktokmall/backend/rpc_gen/kitex_gen/cart"
	order "github.com/tiktokmall/backend/rpc_gen/kitex_gen/order"
)

type ListOderService struct {
	ctx context.Context
} // NewListOderService new ListOderService
func NewListOderService(ctx context.Context) *ListOderService {
	return &ListOderService{ctx: ctx}
}

// Run create note info
func (s *ListOderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	// Finish your business logic.
	orders, err := model.ListOrder(mysql.DB, s.ctx, req.GetUserId())
	var list []*order.Order
	for _, v := range orders {
		var items []*order.OrderItem
		for _, vitem := range v.OrderItem {
			items = append(items, &order.OrderItem{
				Cost: vitem.Cost,
				Item: &cart.CartItem{
					ProductId: vitem.ProductId,
					Quantity:  vitem.Quantity,
				},
			})
		}
		o := &order.Order{
			OrderId:      v.OrderId,
			UserId:       v.UserId,
			UserCurrency: v.UserCurrency,
			Email:        v.Consignee.Email,
			CreatedAt:    int32(v.CreatedAt.Unix()),
			Address: &order.Address{
				Country:       v.Consignee.Country,
				City:          v.Consignee.City,
				StreetAddress: v.Consignee.StreetAddress,
				ZipCode:       v.Consignee.ZipCode,
			},
			Items: items,
		}
		list = append(list, o)
	}
	resp = &order.ListOrderResp{
		Orders: list,
	}
	return
}
