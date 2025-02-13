package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/tiktokmall/backend/app/cart/biz/dal/mysql"
	"github.com/tiktokmall/backend/app/cart/biz/model"
	cart "github.com/tiktokmall/backend/rpc_gen/kitex_gen/cart"
)

type ChangeQtyService struct {
	ctx context.Context
} // NewChangeQtyService new ChangeQtyService
func NewChangeQtyService(ctx context.Context) *ChangeQtyService {
	return &ChangeQtyService{ctx: ctx}
}

// Run create note info
func (s *ChangeQtyService) Run(req *cart.ChangeQtyReq) (resp *cart.ChangeQtyResp, err error) {
	if req.UserId == 0 {
		return nil, kerrors.NewBizStatusError(40004, "user_id is required")
	}
	if req.Item.ProductId == 0 {
		return nil, kerrors.NewBizStatusError(40004, "product_id is required")
	}

	cartItem := &model.Cart{
		UserId:    req.UserId,
		ProductId: req.Item.ProductId,
		Qty:       uint32(req.Item.Quantity),
	}

	err = model.ChangeQty(s.ctx, mysql.DB, cartItem)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50000, err.Error())
	}

	return &cart.ChangeQtyResp{}, nil
}
