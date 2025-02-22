package service

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/tiktokmall/backend/app/merchant/biz/dal/mysql"
	"github.com/tiktokmall/backend/app/merchant/biz/model"
	merchant "github.com/tiktokmall/backend/rpc_gen/kitex_gen/merchant"
)

type AddMerchantService struct {
	ctx context.Context
} // NewAddMerchantService new AddMerchantService
func NewAddMerchantService(ctx context.Context) *AddMerchantService {
	return &AddMerchantService{ctx: ctx}
}

// Run create note info
func (s *AddMerchantService) Run(req *merchant.AddMerchantReq) (resp *merchant.AddMerchantResp, err error) {
	// Finish your business logic.
	if req.Uid <= 0 {
		return nil, kerrors.NewBizStatusError(2004001, "uid is invalid")
	}
	m := model.Merchant{
		UserId:   int(req.Uid),
		Username: req.Username,
		Email:    req.Email,
		Base: model.Base{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	err = model.NewMerchantQuery(s.ctx, mysql.DB).InsertOne(&m)
	if err != nil {
		return nil, err
	}
	resp = &merchant.AddMerchantResp{
		Mid:      int64(m.ID),
		Uid:      req.Uid,
		Username: m.Username,
		Email:    m.Email,
	}
	return
}
