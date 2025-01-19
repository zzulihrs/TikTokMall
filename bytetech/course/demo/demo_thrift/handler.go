package main

import (
	"bytetech/course/demo/demo_thrift/biz/service"
	api "bytetech/course/demo/demo_thrift/kitex_gen/api"
	"context"
)

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// EchoMethod implements the EchoImpl interface.
func (s *EchoImpl) EchoMethod(ctx context.Context, req *api.EchoReq) (resp *api.EchoResp, err error) {
	resp, err = service.NewEchoMethodService(ctx).Run(req)

	return resp, err
}
