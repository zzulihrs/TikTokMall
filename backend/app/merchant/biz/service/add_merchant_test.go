package service

import (
	"context"
	"testing"

	merchant "github.com/tiktokmall/backend/rpc_gen/kitex_gen/merchant"
)

func TestAddMerchant_Run(t *testing.T) {
	ctx := context.Background()
	s := NewAddMerchantService(ctx)
	// init req and assert value

	req := &merchant.AddMerchantReq{
		Uid:      2,
		Username: "test",
		Email:    "EMA",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
