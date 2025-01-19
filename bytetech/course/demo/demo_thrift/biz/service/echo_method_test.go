package service

import (
	api "bytetech/course/demo/demo_thrift/kitex_gen/api"
	"context"
	"testing"
)

func TestEchoMethod_Run(t *testing.T) {
	ctx := context.Background()
	s := NewEchoMethodService(ctx)
	// init req and assert value

	req := &api.EchoReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
