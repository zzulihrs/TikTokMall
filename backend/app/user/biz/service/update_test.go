package service

import (
	"context"
	"testing"
	user "github.com/tiktokmall/backend/rpc_gen/kitex_gen/user"
)

func TestUpdate_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateService(ctx)
	// init req and assert value

	req := &user.UpdateUserReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
