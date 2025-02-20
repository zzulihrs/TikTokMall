package service

import (
	"context"
	order "github.com/tiktokmall/backend/rpc_gen/kitex_gen/order"
	"testing"
)

func TestChangeOrderStatus_Run(t *testing.T) {
	ctx := context.Background()
	s := NewChangeOrderStatusService(ctx)
	// init req and assert value

	req := &order.ChangeOrderStatusReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
