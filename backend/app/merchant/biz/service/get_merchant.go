package service

import (
	"context"

	"github.com/tiktokmall/backend/app/merchant/biz/dal/mysql"
	"github.com/tiktokmall/backend/app/merchant/biz/model"
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

	m, err := model.NewMerchantQuery(s.ctx, mysql.DB).GetByUid(int(req.GetId()))
	if err != nil {
		return nil, err
	}

	resp = &merchant.GetMerchantResp{
		Id:       int64(m.ID),
		Username: m.Username,
		Email:    m.Email,
	}
	return
}
