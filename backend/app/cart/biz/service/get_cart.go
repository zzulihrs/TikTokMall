package service

import (
	"context"

	"github.com/tiktokmall/backend/app/cart/biz/dal/mysql"
	"github.com/tiktokmall/backend/app/cart/biz/model"
	cart "github.com/tiktokmall/backend/rpc_gen/kitex_gen/cart"

	"github.com/cloudwego/kitex/pkg/kerrors"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	// Finish your business logic.
	if req.UserId == 0 {
		return nil, kerrors.NewBizStatusError(40004, "user_id is required")
	}
	carts, err := model.GetCartByUserId(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50000, err.Error())
	}
	items := make([]*cart.CartItem, len(carts))
	for i, v := range carts {
		items[i] = &cart.CartItem{
			ProductId: v.ProductId,
			Quantity:  v.Qty,
		}
	}
	return &cart.GetCartResp{
		Cart: &cart.Cart{
			UserId: req.GetUserId(),
			Items:  items,
		},
	}, nil
}
