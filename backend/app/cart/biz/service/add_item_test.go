package service

import (
	"context"
	"testing"

	cart "github.com/tiktokmall/backend/rpc_gen/kitex_gen/cart"
)

func TestAddItem_Run(t *testing.T) {
	ctx := context.Background()
	s := NewAddItemService(ctx)
	// init req and assert value

	req := &cart.AddItemReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
