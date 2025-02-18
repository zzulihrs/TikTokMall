package service

import (
	"context"

	"github.com/tiktokmall/backend/app/product/biz/dal/mysql"
	"github.com/tiktokmall/backend/app/product/biz/model"
	product "github.com/tiktokmall/backend/rpc_gen/kitex_gen/product"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// Finish your business logic.
	categoryQuery := model.NewCategoryQuery(s.ctx, mysql.DB)
	cs, err := categoryQuery.GetProductByCategoryName(req.CategoryName)
	if err != nil {
		return nil, err
	}
	// TODO: 添加商家信息
	merchantSet := make(map[int]int)
	merchantIDs := []int{}

	resp = &product.ListProductsResp{}
	for _, c := range cs {
		for _, p := range c.Products {
			resp.Products = append(resp.Products, &product.Product{
				Id:          uint32(p.ID),
				Name:        p.Name,
				Price:       p.Price,
				Description: p.Description,
				Picture:     p.Picture,
				Stock:       uint32(p.Stock),
				OwnerId:     uint32(p.MerchantID),
			})
			if merchantSet[p.MerchantID] == 0 {
				merchantSet[p.MerchantID] = len(merchantIDs) + 1
				merchantIDs = append(merchantIDs, p.MerchantID)
			}
		}
	}

	merchants, err := model.NewMerchantQuery(s.ctx, mysql.DB).GetManyById(merchantIDs)
	if err != nil {
		return nil, err
	}

	for _, p := range resp.Products {
		mIndex := merchantSet[int(p.OwnerId)] - 1
		p.OwnerName = merchants[mIndex].Username
	}

	return resp, nil
}
