package merchant

import (
	"context"
	merchant "github.com/tiktokmall/backend/rpc_gen/kitex_gen/merchant"

	"github.com/tiktokmall/backend/rpc_gen/kitex_gen/merchant/merchantservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() merchantservice.Client
	Service() string
	GetMerchant(ctx context.Context, Req *merchant.GetMerchantReq, callOptions ...callopt.Option) (r *merchant.GetMerchantResp, err error)
	AddProduct(ctx context.Context, Req *merchant.AddProductReq, callOptions ...callopt.Option) (r *merchant.AddProductResp, err error)
	DeleteProduct(ctx context.Context, Req *merchant.DeleteProductReq, callOptions ...callopt.Option) (r *merchant.DeleteProductResp, err error)
	UpdateProduct(ctx context.Context, Req *merchant.UpdateProductReq, callOptions ...callopt.Option) (r *merchant.UpdateProductResp, err error)
	SearchProducts(ctx context.Context, Req *merchant.SearchProductsReq, callOptions ...callopt.Option) (r *merchant.SearchProductsResp, err error)
	ProductDetail(ctx context.Context, Req *merchant.ProductDetailReq, callOptions ...callopt.Option) (r *merchant.ProductDetailResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := merchantservice.NewClient(dstService, opts...)
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
	kitexClient merchantservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() merchantservice.Client {
	return c.kitexClient
}

func (c *clientImpl) GetMerchant(ctx context.Context, Req *merchant.GetMerchantReq, callOptions ...callopt.Option) (r *merchant.GetMerchantResp, err error) {
	return c.kitexClient.GetMerchant(ctx, Req, callOptions...)
}

func (c *clientImpl) AddProduct(ctx context.Context, Req *merchant.AddProductReq, callOptions ...callopt.Option) (r *merchant.AddProductResp, err error) {
	return c.kitexClient.AddProduct(ctx, Req, callOptions...)
}

func (c *clientImpl) DeleteProduct(ctx context.Context, Req *merchant.DeleteProductReq, callOptions ...callopt.Option) (r *merchant.DeleteProductResp, err error) {
	return c.kitexClient.DeleteProduct(ctx, Req, callOptions...)
}

func (c *clientImpl) UpdateProduct(ctx context.Context, Req *merchant.UpdateProductReq, callOptions ...callopt.Option) (r *merchant.UpdateProductResp, err error) {
	return c.kitexClient.UpdateProduct(ctx, Req, callOptions...)
}

func (c *clientImpl) SearchProducts(ctx context.Context, Req *merchant.SearchProductsReq, callOptions ...callopt.Option) (r *merchant.SearchProductsResp, err error) {
	return c.kitexClient.SearchProducts(ctx, Req, callOptions...)
}

func (c *clientImpl) ProductDetail(ctx context.Context, Req *merchant.ProductDetailReq, callOptions ...callopt.Option) (r *merchant.ProductDetailResp, err error) {
	return c.kitexClient.ProductDetail(ctx, Req, callOptions...)
}
