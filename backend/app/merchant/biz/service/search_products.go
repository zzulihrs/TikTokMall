package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/tiktokmall/backend/app/merchant/biz/dal/mysql"
	"github.com/tiktokmall/backend/app/merchant/biz/model"
	merchant "github.com/tiktokmall/backend/rpc_gen/kitex_gen/merchant"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *merchant.SearchProductsReq) (resp *merchant.SearchProductsResp, err error) {
	// Finish your business logic.
	// TODO:
	// 1. 检查参数
	if req.GetMerchantId() <= 0 {
		return nil, kerrors.NewBizStatusError(2004001, "merchant id must be > 0")
	}
	// 2. 检查商户
	_, err = model.NewMerchantQuery(s.ctx, mysql.DB).GetById(int(req.GetMerchantId()))
	if err != nil {
		return nil, fmt.Errorf("merchant [%v] not found, err:%w", req.GetMerchantId(), err)
	}
	// 3. 规整条件
	conditions := []string{}
	if req.GetName() != "" {
		conditions = append(conditions, fmt.Sprintf("name LIKE '%%%v%%'", req.GetName()))
	}
	if req.GetCategoryId() > 0 {
		conditions = append(conditions, fmt.Sprintf("category_id = %v", req.GetCategoryId()))
	}
	condition := strings.Join(conditions, " AND ")
	// 4. 搜索商品
	_, err = model.NewProductQuery(s.ctx, mysql.DB).GetProductListByCondition(condition)
	return
}
