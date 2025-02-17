package service

import (
	"context"
	merchant "github.com/tiktokmall/backend/rpc_gen/kitex_gen/merchant"
)

type DeleteProductService struct {
	ctx context.Context
} // NewDeleteProductService new DeleteProductService
func NewDeleteProductService(ctx context.Context) *DeleteProductService {
	return &DeleteProductService{ctx: ctx}
}

// Run create note info
func (s *DeleteProductService) Run(req *merchant.DeleteProductReq) (resp *merchant.DeleteProductResp, err error) {
	// Finish your business logic.

	return
}
