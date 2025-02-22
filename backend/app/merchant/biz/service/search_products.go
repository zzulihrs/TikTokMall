package service

import (
	"context"
	"fmt"
	"log"
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
	// TODO: 列表条件查询
	// 1. 检查参数
	// log.Printf("search products, req: %+v", req)
	if req.GetMerchantId() <= 0 {
		return nil, kerrors.NewBizStatusError(2004001, "merchant id must be > 0")
	}
	if req.GetMaxPrice() < req.GetMinPrice() {
		return nil, kerrors.NewBizStatusError(2004001, "max price must be >= min price")
	}
	// 2. 规整条件

	conditions := []string{fmt.Sprintf("merchant_id = %v", req.GetMerchantId())}
	if req.GetName() != "" {
		conditions = append(conditions, fmt.Sprintf("name LIKE '%%%v%%'", req.GetName()))
	}
	// if req.GetCategoryId() > 0 {
	// 	conditions = append(conditions, fmt.Sprintf("category_id = %v", req.GetCategoryId()))
	// }
	if req.GetMinPrice() > 0 {
		conditions = append(conditions, fmt.Sprintf("price >= %v", req.GetMinPrice()))
	}
	if req.GetMaxPrice() > 0 {
		conditions = append(conditions, fmt.Sprintf("price <= %v", req.GetMaxPrice()))
	}
	if req.GetMaxStock() > 0 {
		conditions = append(conditions, fmt.Sprintf("stock <= %v", req.GetMaxStock()))
	}
	// 3. 拼接条件
	condition := strings.Join(conditions, " AND ")
	log.Printf("condition: %v", condition)
	// 4. 分页
	pageNo := req.GetPageNo()
	if pageNo == 0 {
		pageNo = 1
	}
	pageSize := req.GetPageSize()
	if pageSize == 0 {
		pageSize = 10
	}
	// 5. 搜索商品
	products, count, err := model.NewProductQuery(s.ctx, mysql.DB).GetProductListByCondition(condition, int(pageNo), int(pageSize))
	rps := make([]*merchant.MerchantProductSimpleInfo, len(products))
	for i, p := range products {
		rps[i] = &merchant.MerchantProductSimpleInfo{
			Id:          int64(p.ID),
			Name:        p.Name,
			Description: p.Description,
			ImgUrl:      p.Picture,
			Price:       p.Price,
			Stock:       int32(p.Stock),
		}
	}

	resp = &merchant.SearchProductsResp{
		Products: rps,
		Count:    count,
	}
	return
}
