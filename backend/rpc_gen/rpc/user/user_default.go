package user

import (
	"context"
	user "github.com/tiktokmall/backend/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Register(ctx context.Context, req *user.RegisterReq, callOptions ...callopt.Option) (resp *user.RegisterResp, err error) {
	resp, err = defaultClient.Register(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Register call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Login(ctx context.Context, req *user.LoginReq, callOptions ...callopt.Option) (resp *user.LoginResp, err error) {
	resp, err = defaultClient.Login(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Login call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Update(ctx context.Context, req *user.UpdateUserReq, callOptions ...callopt.Option) (resp *user.UpdateUserResp, err error) {
	resp, err = defaultClient.Update(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Update call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Delete(ctx context.Context, req *user.DeleteUserReq, callOptions ...callopt.Option) (resp *user.DeleteUserResp, err error) {
	resp, err = defaultClient.Delete(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Delete call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Query(ctx context.Context, req *user.QueryUserReq, callOptions ...callopt.Option) (resp *user.QueryUserResp, err error) {
	resp, err = defaultClient.Query(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Query call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
