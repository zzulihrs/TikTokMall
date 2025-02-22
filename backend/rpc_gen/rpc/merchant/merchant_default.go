package merchant

import (
	"context"
	merchant "github.com/tiktokmall/backend/rpc_gen/kitex_gen/merchant"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func GetMerchant(ctx context.Context, req *merchant.GetMerchantReq, callOptions ...callopt.Option) (resp *merchant.GetMerchantResp, err error) {
	resp, err = defaultClient.GetMerchant(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetMerchant call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func AddMerchant(ctx context.Context, req *merchant.AddMerchantReq, callOptions ...callopt.Option) (resp *merchant.AddMerchantResp, err error) {
	resp, err = defaultClient.AddMerchant(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "AddMerchant call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func AddProduct(ctx context.Context, req *merchant.AddProductReq, callOptions ...callopt.Option) (resp *merchant.AddProductResp, err error) {
	resp, err = defaultClient.AddProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "AddProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteProduct(ctx context.Context, req *merchant.DeleteProductReq, callOptions ...callopt.Option) (resp *merchant.DeleteProductResp, err error) {
	resp, err = defaultClient.DeleteProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateProduct(ctx context.Context, req *merchant.UpdateProductReq, callOptions ...callopt.Option) (resp *merchant.UpdateProductResp, err error) {
	resp, err = defaultClient.UpdateProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func SearchProducts(ctx context.Context, req *merchant.SearchProductsReq, callOptions ...callopt.Option) (resp *merchant.SearchProductsResp, err error) {
	resp, err = defaultClient.SearchProducts(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "SearchProducts call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ProductDetail(ctx context.Context, req *merchant.ProductDetailReq, callOptions ...callopt.Option) (resp *merchant.ProductDetailResp, err error) {
	resp, err = defaultClient.ProductDetail(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ProductDetail call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
