package service

import (
	pbapi "bytetech/course/demo/demo_proto/kitex_gen/pbapi"
	"context"
)

type EchoMethodService struct {
	ctx context.Context
} // NewEchoMethodService new EchoMethodService
func NewEchoMethodService(ctx context.Context) *EchoMethodService {
	return &EchoMethodService{ctx: ctx}
}

// Run create note info
func (s *EchoMethodService) Run(req *pbapi.EchoRequest) (resp *pbapi.EchoResponse, err error) {
	// Finish your business logic.

	return &pbapi.EchoResponse{
		Message: req.Message,
	}, nil
}
