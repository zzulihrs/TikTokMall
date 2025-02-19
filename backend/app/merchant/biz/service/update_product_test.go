package service

import (
	"context"
	"testing"

	merchant "github.com/tiktokmall/backend/rpc_gen/kitex_gen/merchant"
)

func TestUpdateProduct_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateProductService(ctx)
	// init req and assert value

	req := &merchant.UpdateProductReq{
		MerchantId: 1,
		Product: &merchant.MerchantProductDetailInfo{
			Id:   9,
			Name: "test",
		},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
