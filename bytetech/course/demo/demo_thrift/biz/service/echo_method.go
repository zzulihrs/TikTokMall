package service

import (
	api "bytetech/course/demo/demo_thrift/kitex_gen/api"
	"context"
)

type EchoMethodService struct {
	ctx context.Context
} // NewEchoMethodService new EchoMethodService
func NewEchoMethodService(ctx context.Context) *EchoMethodService {
	return &EchoMethodService{ctx: ctx}
}

// Run create note info
func (s *EchoMethodService) Run(req *api.EchoReq) (resp *api.EchoResp, err error) {
	// Finish your business logic.

	return &api.EchoResp{Message: req.Message}, nil
}
