package service

import (
	"context"
	"testing"
	payment "github.com/tiktokmall/backend/rpc_gen/kitex_gen/payment"
)

func TestAlipay_Run(t *testing.T) {
	ctx := context.Background()
	s := NewAlipayService(ctx)
	// init req and assert value

	req := &payment.AlipayReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
