package service

import (
	"context"
	"strings"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/tiktokmall/backend/app/merchant/biz/dal/mysql"
	"github.com/tiktokmall/backend/app/merchant/biz/model"
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
	// 1. 检查参数
	if req.GetMerchantId() <= 0 {
		return nil, kerrors.NewBizStatusError(2004001, "merchant id must be > 0")
	}
	if req.GetPid() <= 0 {
		return nil, kerrors.NewBizStatusError(2004001, "product id must be > 0")
	}
	product, err := model.NewProductQuery(s.ctx, mysql.DB).GetById(int(req.GetPid()))
	if product.MerchantID != int(req.GetMerchantId()) {
		return nil, kerrors.NewBizStatusError(2004001, "merchant id not match")
	}
	categories := make([]*merchant.Category, len(product.Categories))
	for i, c := range product.Categories {
		categories[i] = &merchant.Category{
			Id:          int64(c.ID),
			Name:        c.Name,
			Description: c.Description,
		}
	}

	resp = &merchant.ProductDetailResp{
		Product: &merchant.MerchantProductDetailInfo{
			Id:          int64(product.ID),
			Name:        product.Name,
			Description: product.Description,
			ImgUrl:      product.Picture,
			Price:       product.Price,
			SliderImgs:  strings.Split(product.SliderImgs, "\n"),
			Stock:       int32(product.Stock),
			Category:    categories,
		},
	}

	return
}
