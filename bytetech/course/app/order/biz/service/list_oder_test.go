package service

import (
	"context"
	"testing"
	order "bytetech/course/rpc_gen/kitex_gen/order"
)

func TestListOder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewListOderService(ctx)
	// init req and assert value

	req := &order.ListOrderReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
