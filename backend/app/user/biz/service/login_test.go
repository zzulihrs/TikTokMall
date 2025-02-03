package service

import (
	"context"
	"testing"

	user "github.com/tiktokmall/backend/rpc_gen/kitex_gen/user"
)

func TestLogin_Run(t *testing.T) {
	ctx := context.Background()
	s := NewLoginService(ctx)
	// init req and assert value

	req := &user.LoginReq{
		Email:    "test@test.com",
		Password: "123456",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
