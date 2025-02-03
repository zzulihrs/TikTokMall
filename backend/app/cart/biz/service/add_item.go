package service

import (
	"context"

	"github.com/tiktokmall/backend/app/cart/biz/dal/mysql"
	"github.com/tiktokmall/backend/app/cart/biz/model"
	"github.com/tiktokmall/backend/app/cart/infra/rpc"
	cart "github.com/tiktokmall/backend/rpc_gen/kitex_gen/cart"
	"github.com/tiktokmall/backend/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/kitex/pkg/kerrors"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	// Finish your business logic.

	if req.UserId == 0 {
		return nil, kerrors.NewBizStatusError(40004, "user_id is required")
	}
	if req.Item.ProductId == 0 {
		return nil, kerrors.NewBizStatusError(40004, "product_id is required")
	}

	productResp, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: req.Item.ProductId})
	if err != nil {
		return nil, err
	}
	if productResp == nil || productResp.Product.Id == 0 {
		return nil, kerrors.NewBizStatusError(40004, "product not found")
	}

	err = model.AddCart(mysql.DB, s.ctx, &model.Cart{
		UserId:    req.UserId,
		ProductId: req.Item.ProductId,
		Qty:       uint32(req.Item.Quantity),
	})
	if err != nil {
		return nil, kerrors.NewBizStatusError(50000, err.Error())
	}

	return &cart.AddItemResp{}, nil
}
