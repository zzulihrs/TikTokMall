package service

import (
	"context"
	"fmt"
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

func TestCrypt(t *testing.T) {
	// Test case 1: Normal password
	password1 := "123456"
	hashedPassword1, _ := crypt(password1)
	// assert.NoError(t, err1)
	// assert.NotEmpty(t, hashedPassword1)
	fmt.Println(hashedPassword1)
}
