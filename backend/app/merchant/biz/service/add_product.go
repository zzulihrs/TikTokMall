package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/tiktokmall/backend/app/merchant/biz/dal/mysql"
	"github.com/tiktokmall/backend/app/merchant/biz/model"
	merchant "github.com/tiktokmall/backend/rpc_gen/kitex_gen/merchant"
)

type AddProductService struct {
	ctx context.Context
} // NewAddProductService new AddProductService
func NewAddProductService(ctx context.Context) *AddProductService {
	return &AddProductService{ctx: ctx}
}

// Run create note info
func (s *AddProductService) Run(req *merchant.AddProductReq) (resp *merchant.AddProductResp, err error) {
	// Finish your business logic.

	// 1. 检查参数
	if err = checkAddProductReq(req); err != nil {
		return nil, kerrors.NewGRPCBizStatusError(2004001, "product is invalid")
	}
	// 2. 检查商户
	_, err = model.NewMerchantQuery(s.ctx, mysql.DB).GetById(int(req.GetMerchantId()))
	if err != nil {
		return nil, fmt.Errorf("merchant [%v] not found, err:%w", req.GetMerchantId(), err)
	}
	// 3. 检查类别,只看id
	reqCategories := req.GetProduct().GetCategory()
	reqCsIds := []int64{}
	for _, c := range reqCategories {
		reqCsIds = append(reqCsIds, c.GetId())
	}
	categories, err := model.NewCategoryQuery(s.ctx, mysql.DB).GetManyById(reqCsIds)
	if err != nil {
		return nil, fmt.Errorf("cannot find all categories [%v], err:%w", reqCsIds, err)
	}
	// 4. 插入商品 与 类别关系
	newProduct := &model.Product{
		Name:        req.GetProduct().GetName(),
		Description: req.GetProduct().GetDescription(),
		Picture:     req.GetProduct().GetImgUrl(),
		Price:       req.GetProduct().GetPrice(),
		SliderImgs:  strings.Join(req.GetProduct().GetSliderImgs(), "\n"),
		Stock:       int32(req.GetProduct().GetStock()),
		MerchantID:  int(req.GetMerchantId()),
		Categories:  categories,
		Base: model.Base{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	err = model.NewProductQuery(s.ctx, mysql.DB).InsertMany([]model.Product{*newProduct})
	if err != nil {
		return nil, fmt.Errorf("insert product failed, err:%w", err)
	}

	return
}

func checkAddProductReq(req *merchant.AddProductReq) error {
	if req.GetMerchantId() <= 0 {
		return kerrors.NewBizStatusError(2004001, "merchant id must be > 0")
	}
	if req.GetProduct() == nil {
		return kerrors.NewBizStatusError(2004001, "product is required")
	}
	reqProduct := req.GetProduct()
	if reqProduct.GetName() == "" {
		return kerrors.NewBizStatusError(2004001, "product name is required")
	}
	if reqProduct.GetPrice() < 0 {
		return kerrors.NewBizStatusError(2004001, "product price must be >= 0")
	}
	if reqProduct.GetStock() < 0 {
		return kerrors.NewBizStatusError(2004001, "product stock must be >= 0")
	}
	reqCategory := reqProduct.GetCategory()
	if reqCategory == nil {
		return kerrors.NewBizStatusError(2004001, "product category is required")
	}
	return nil
}
