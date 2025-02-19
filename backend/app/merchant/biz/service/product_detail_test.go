package service

import (
	"context"
	"testing"

	merchant "github.com/tiktokmall/backend/rpc_gen/kitex_gen/merchant"
)

func TestProductDetail_Run(t *testing.T) {
	ctx := context.Background()
	s := NewProductDetailService(ctx)
	// init req and assert value

	req := &merchant.ProductDetailReq{
		MerchantId: 1,
		Pid:        7,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
