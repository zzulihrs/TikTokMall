package service

import (
	"context"
	"testing"
	pbapi "bytetech/course/demo/demo_proto/kitex_gen/pbapi"
)

func TestEchoMethod_Run(t *testing.T) {
	ctx := context.Background()
	s := NewEchoMethodService(ctx)
	// init req and assert value

	req := &pbapi.EchoRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
