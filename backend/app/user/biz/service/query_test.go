package service

import (
	"context"
	"testing"
	user "github.com/tiktokmall/backend/rpc_gen/kitex_gen/user"
)

func TestQuery_Run(t *testing.T) {
	ctx := context.Background()
	s := NewQueryService(ctx)
	// init req and assert value

	req := &user.QueryUserReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
