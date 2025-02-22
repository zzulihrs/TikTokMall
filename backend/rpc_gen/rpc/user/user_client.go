package user

import (
	"context"
	user "github.com/tiktokmall/backend/rpc_gen/kitex_gen/user"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/tiktokmall/backend/rpc_gen/kitex_gen/user/userservice"
)

type RPCClient interface {
	KitexClient() userservice.Client
	Service() string
	Register(ctx context.Context, Req *user.RegisterReq, callOptions ...callopt.Option) (r *user.RegisterResp, err error)
	Login(ctx context.Context, Req *user.LoginReq, callOptions ...callopt.Option) (r *user.LoginResp, err error)
	Update(ctx context.Context, Req *user.UpdateUserReq, callOptions ...callopt.Option) (r *user.UpdateUserResp, err error)
	Delete(ctx context.Context, Req *user.DeleteUserReq, callOptions ...callopt.Option) (r *user.DeleteUserResp, err error)
	Query(ctx context.Context, Req *user.QueryUserReq, callOptions ...callopt.Option) (r *user.QueryUserResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := userservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient userservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() userservice.Client {
	return c.kitexClient
}

func (c *clientImpl) Register(ctx context.Context, Req *user.RegisterReq, callOptions ...callopt.Option) (r *user.RegisterResp, err error) {
	return c.kitexClient.Register(ctx, Req, callOptions...)
}

func (c *clientImpl) Login(ctx context.Context, Req *user.LoginReq, callOptions ...callopt.Option) (r *user.LoginResp, err error) {
	return c.kitexClient.Login(ctx, Req, callOptions...)
}

func (c *clientImpl) Update(ctx context.Context, Req *user.UpdateUserReq, callOptions ...callopt.Option) (r *user.UpdateUserResp, err error) {
	return c.kitexClient.Update(ctx, Req, callOptions...)
}

func (c *clientImpl) Delete(ctx context.Context, Req *user.DeleteUserReq, callOptions ...callopt.Option) (r *user.DeleteUserResp, err error) {
	return c.kitexClient.Delete(ctx, Req, callOptions...)
}

func (c *clientImpl) Query(ctx context.Context, Req *user.QueryUserReq, callOptions ...callopt.Option) (r *user.QueryUserResp, err error) {
	return c.kitexClient.Query(ctx, Req, callOptions...)
}
