package service

import (
	"context"
	"testing"
	merchant "github.com/tiktokmall/backend/rpc_gen/kitex_gen/merchant"
)

func TestAddProduct_Run(t *testing.T) {
	ctx := context.Background()
	s := NewAddProductService(ctx)
	// init req and assert value

	req := &merchant.AddProductReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
