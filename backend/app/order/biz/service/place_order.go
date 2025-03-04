package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"github.com/tiktokmall/backend/app/order/biz/dal/mysql"
	"github.com/tiktokmall/backend/app/order/biz/model"
	"github.com/tiktokmall/backend/app/order/infra/mq"
	order "github.com/tiktokmall/backend/rpc_gen/kitex_gen/order"
	"gorm.io/gorm"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	// Finish your business logic.
	if len(req.Items) == 0 {
		err = kerrors.NewBizStatusError(500001, "items is empty")
		return
	}
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		orderId, _ := uuid.NewUUID()

		o := &model.Order{
			OrderId:      orderId.String(),
			UserId:       req.UserId,
			UserCurrency: req.UserCurrency,
			Consignee: model.Consignee{
				Email: req.Email,
			},
			OrderStatus: 0, // 0: created
		}
		if req.Address != nil {
			a := req.Address
			o.Consignee.StreetAddress = a.StreetAddress
			o.Consignee.City = a.City
			o.Consignee.State = a.State
			o.Consignee.Country = a.Country
		}

		if err := tx.Create(o).Error; err != nil {
			return err
		}

		var items []model.OrderItem
		for _, v := range req.Items {
			items = append(items, model.OrderItem{
				OrderIdRefer: orderId.String(),
				ProductId:    v.Item.ProductId,
				Quantity:     int32(v.Item.Quantity),
				Cost:         v.Cost,
			})
		}

		if err := tx.Create(items).Error; err != nil {
			return err
		}

		// 使用nats延迟队列，延迟30分钟执行，id为orderId
		msg := &nats.Msg{
			Subject: "ORDER.CREATE",
			Data:    []byte(orderId.String()),
			Header: nats.Header{
				"Nats-Delay": []string{"10"}, // 单位：秒
			},
		}

		if err := mq.Nc.PublishMsg(msg); err != nil {
			return err
		}

		resp = &order.PlaceOrderResp{
			Order: &order.OrderResult{
				OrderId: orderId.String(),
			},
		}

		return nil
	})

	return
}
