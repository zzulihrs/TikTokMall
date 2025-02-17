package service

import (
	"context"
	merchant "github.com/tiktokmall/backend/rpc_gen/kitex_gen/merchant"
)

type ProductDetailService struct {
	ctx context.Context
} // NewProductDetailService new ProductDetailService
func NewProductDetailService(ctx context.Context) *ProductDetailService {
	return &ProductDetailService{ctx: ctx}
}

// Run create note info
func (s *ProductDetailService) Run(req *merchant.ProductDetailReq) (resp *merchant.ProductDetailResp, err error) {
	// Finish your business logic.

	return
}
