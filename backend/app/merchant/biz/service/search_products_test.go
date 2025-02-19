package service

import (
	"context"
	"testing"

	merchant "github.com/tiktokmall/backend/rpc_gen/kitex_gen/merchant"
)

func TestSearchProducts_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSearchProductsService(ctx)
	// init req and assert value

	req := &merchant.SearchProductsReq{
		MerchantId: 1,
		Name:       "t",
		MinPrice:   5.0,
		MaxPrice:   10.0,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
