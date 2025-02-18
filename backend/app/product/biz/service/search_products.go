package service

import (
	"context"

	"github.com/tiktokmall/backend/app/product/biz/dal/mysql"
	"github.com/tiktokmall/backend/app/product/biz/model"
	product "github.com/tiktokmall/backend/rpc_gen/kitex_gen/product"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	// Finish your business logic.
	pq := model.NewProductQuery(s.ctx, mysql.DB)
	ps, err := pq.SearchProducts(req.Query)
	if err != nil {
		return nil, err
	}
	results := []*product.Product{}
	for _, p := range ps {
		results = append(results, &product.Product{
			Id:          uint32(p.ID),
			Name:        p.Name,
			Price:       p.Price,
			Description: p.Description,
			Picture:     p.Picture})
	}

	return &product.SearchProductsResp{Results: results}, nil
}
