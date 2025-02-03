package service

import (
	"context"

	"github.com/tiktokmall/backend/app/cart/biz/dal/mysql"
	"github.com/tiktokmall/backend/app/cart/biz/model"
	cart "github.com/tiktokmall/backend/rpc_gen/kitex_gen/cart"

	"github.com/cloudwego/kitex/pkg/kerrors"
)

type EmptyCartService struct {
	ctx context.Context
} // NewEmptyCartService new EmptyCartService
func NewEmptyCartService(ctx context.Context) *EmptyCartService {
	return &EmptyCartService{ctx: ctx}
}

// Run create note info
func (s *EmptyCartService) Run(req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	// Finish your business logic.
	if req.UserId == 0 {
		return nil, kerrors.NewBizStatusError(40004, "user_id is required")
	}
	err = model.EmptyCart(mysql.DB, s.ctx, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50000, err.Error())
	}
	return &cart.EmptyCartResp{}, nil
}
