package service

import (
	"bytetech/course/app/cart/biz/dal/mysql"
	"bytetech/course/app/cart/biz/model"
	"bytetech/course/app/cart/infra/rpc"
	cart "bytetech/course/rpc_gen/kitex_gen/cart"
	"bytetech/course/rpc_gen/kitex_gen/product"
	"context"

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
