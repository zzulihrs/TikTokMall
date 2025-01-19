package service

import (
	"bytetech/course/app/cart/biz/dal/mysql"
	"bytetech/course/app/cart/biz/model"
	cart "bytetech/course/rpc_gen/kitex_gen/cart"
	"context"

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
	carts, err := model.GetCartByUserId(mysql.DB, s.ctx, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50000, err.Error())
	}
	items := make([]*cart.CartItem, len(carts))
	for i, v := range carts {
		items[i] = &cart.CartItem{
			ProductId: v.ProductId,
			Quantity:  int32(v.Qty),
		}
	}
	return &cart.GetCartResp{
		Cart: &cart.Cart{
			UserId: req.GetUserId(),
			Items:  items,
		},
	}, nil
}
