package service

import (
	"bytetech/course/app/product/biz/dal/mysql"
	"bytetech/course/app/product/biz/model"
	product "bytetech/course/rpc_gen/kitex_gen/product"
	"context"
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
	resp = &product.ListProductsResp{}
	for _, c := range cs {
		for _, p := range c.Products {
			resp.Products = append(resp.Products, &product.Product{
				Id:          uint32(p.ID),
				Name:        p.Name,
				Price:       p.Price,
				Description: p.Description,
				Picture:     p.Picture,
			})
		}
	}
	return resp, nil
}
