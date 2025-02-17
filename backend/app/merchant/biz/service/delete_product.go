package service

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/tiktokmall/backend/app/merchant/biz/dal/mysql"
	"github.com/tiktokmall/backend/app/merchant/biz/model"
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

	// 1. 检查参数
	if req.GetMerchantId() <= 0 {
		return nil, kerrors.NewGRPCBizStatusError(2004001, "merchant id must be > 0")
	}
	if req.GetPid() <= 0 {
		return nil, kerrors.NewGRPCBizStatusError(2004001, "product id must be > 0")
	}

	// 2. 检查商户
	_, err = model.NewMerchantQuery(s.ctx, mysql.DB).GetById(int(req.GetMerchantId()))
	if err != nil {
		return nil, fmt.Errorf("merchant [%v] not found, err:%w", req.GetMerchantId(), err)
	}
	err = model.NewProductQuery(s.ctx, mysql.DB).DeleteByPidAndMid(int(req.GetPid()), int(req.GetMerchantId()))
	if err != nil {
		return nil, fmt.Errorf("delete product failed, err:%w", err)
	}
	return
}
