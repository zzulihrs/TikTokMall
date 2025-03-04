package service

import (
	"context"

	"github.com/tiktokmall/backend/app/product/biz/dal/mysql"
	"github.com/tiktokmall/backend/app/product/biz/dal/redis"
	"github.com/tiktokmall/backend/app/product/biz/model"
	product "github.com/tiktokmall/backend/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/kitex/pkg/kerrors"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	// Finish your business logic.
	if req.Id == 0 {
		return nil, kerrors.NewGRPCBizStatusError(2004001, "product id is required")
	}
	// pq := model.NewCachedProductQuery(model.NewProductQuery(s.ctx, mysql.DB), redis.RedisClient, "shop")
	// p, err := pq.GetById(int(req.Id))
	// if err != nil {
	// 	return nil, err
	// }
	// merchant, err := model.NewMerchantQuery(s.ctx, mysql.DB).GetById(p.MerchantID)
	// if err != nil {
	// 	return nil, err
	// }
	// // 只添加类别名称
	// var categories []string
	// for _, c := range p.Categories {
	// 	categories = append(categories, c.Name)
	// }
	// log.Println("merchant: ", merchant)
	// return &product.GetProductResp{
	// 	Product: &product.Product{
	// 		Id:          uint32(p.ID),
	// 		Picture:     p.Picture,
	// 		Price:       p.Price,
	// 		Description: p.Description,
	// 		Name:        p.Name,
	// 		Stock:       uint32(p.Stock),
	// 		OwnerId:     uint32(merchant.ID),
	// 		OwnerName:   merchant.Username,
	// 		Categories:  categories,
	// 	},
	// }, nil

	p, err := model.NewProductDAO(s.ctx, mysql.DB, redis.RedisClient).GetProductDetailByPid(int64(req.Id))
	if err != nil {
		return nil, err
	}
	return &product.GetProductResp{
		Product: &product.Product{
			Id:          uint32(p.ID),
			Picture:     p.Picture,
			Price:       p.Price,
			Description: p.Description,
			Name:        p.Name,
			Stock:       uint32(p.Stock),
			OwnerId:     uint32(p.OwnerId),
			OwnerName:   p.OwnerName,
			Categories:  p.CategorieNames,
		},
	}, nil
}
