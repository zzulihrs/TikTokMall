package main

import (
	"context"

	"github.com/tiktokmall/backend/app/user/biz/service"
	user "github.com/tiktokmall/backend/rpc_gen/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	resp, err = service.NewRegisterService(ctx).Run(req)

	return resp, err
}

// // Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	resp, err = service.NewLoginService(ctx).Run(req)

	return resp, err
}

// Update implements the UserServiceImpl interface.
func (s *UserServiceImpl) Update(ctx context.Context, req *user.UpdateUserReq) (resp *user.UpdateUserResp, err error) {
	resp, err = service.NewUpdateService(ctx).Run(req)

	return resp, err
}

// Delete implements the UserServiceImpl interface.
func (s *UserServiceImpl) Delete(ctx context.Context, req *user.DeleteUserReq) (resp *user.DeleteUserResp, err error) {
	resp, err = service.NewDeleteService(ctx).Run(req)

	return resp, err
}

// Query implements the UserServiceImpl interface.
func (s *UserServiceImpl) Query(ctx context.Context, req *user.QueryUserReq) (resp *user.QueryUserResp, err error) {
	resp, err = service.NewQueryService(ctx).Run(req)

	return resp, err
}
