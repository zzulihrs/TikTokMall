package service

import (
	"context"
	merchant "github.com/tiktokmall/backend/rpc_gen/kitex_gen/merchant"
)

type GetMerchantService struct {
	ctx context.Context
} // NewGetMerchantService new GetMerchantService
func NewGetMerchantService(ctx context.Context) *GetMerchantService {
	return &GetMerchantService{ctx: ctx}
}

// Run create note info
func (s *GetMerchantService) Run(req *merchant.GetMerchantReq) (resp *merchant.GetMerchantResp, err error) {
	// Finish your business logic.

	return
}
