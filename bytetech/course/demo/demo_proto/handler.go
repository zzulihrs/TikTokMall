package main

import (
	pbapi "bytetech/course/demo/demo_proto/kitex_gen/pbapi"
	"context"
	"bytetech/course/demo/demo_proto/biz/service"
)

// EchoServiceImpl implements the last service interface defined in the IDL.
type EchoServiceImpl struct{}

// EchoMethod implements the EchoServiceImpl interface.
func (s *EchoServiceImpl) EchoMethod(ctx context.Context, req *pbapi.EchoRequest) (resp *pbapi.EchoResponse, err error) {
	resp, err = service.NewEchoMethodService(ctx).Run(req)

	return resp, err
}
