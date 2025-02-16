package service

import (
	"context"
	"log"

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
	log.Printf("req: %v\n", req)
	ps, err := pq.SearchProducts(req.Query)
	log.Printf("ps: %#v\n", ps)
	if err != nil {
		return nil, err
	}

	results := []*product.Product{}
	// TODO: 添加商家信息
	merchantSet := make(map[int]int)
	merchantIDs := []int{}
	for _, p := range ps {
		log.Printf("p: %#v\n", p)
		results = append(results, &product.Product{
			Id:          uint32(p.ID),
			Name:        p.Name,
			Price:       p.Price,
			Description: p.Description,
			Picture:     p.Picture,
			Stock:       uint32(p.Stock),
			Categories:  []string{},
			OwnerId:     uint32(p.MerchantID),
		})
		if merchantSet[p.MerchantID] == 0 {
			merchantSet[p.MerchantID] = len(merchantIDs) + 1
			merchantIDs = append(merchantIDs, p.MerchantID)
		}
	}

	merchants, err := model.NewMerchantQuery(s.ctx, mysql.DB).GetManyById(merchantIDs)
	log.Println(err)
	if err != nil {
		return nil, err
	}
	log.Printf("merchants: %#v\n", merchants)

	for i, p := range results {
		mIndex := merchantSet[int(p.OwnerId)] - 1
		results[i].OwnerName = merchants[mIndex].Username
	}
	log.Printf("results: %#v\n", results)

	return &product.SearchProductsResp{Results: results}, nil
}
