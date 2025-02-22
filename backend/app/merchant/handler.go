package main

import (
	"context"
	merchant "github.com/tiktokmall/backend/rpc_gen/kitex_gen/merchant"
	"github.com/tiktokmall/backend/app/merchant/biz/service"
)

// MerchantServiceImpl implements the last service interface defined in the IDL.
type MerchantServiceImpl struct{}

// GetMerchant implements the MerchantServiceImpl interface.
func (s *MerchantServiceImpl) GetMerchant(ctx context.Context, req *merchant.GetMerchantReq) (resp *merchant.GetMerchantResp, err error) {
	resp, err = service.NewGetMerchantService(ctx).Run(req)

	return resp, err
}

// AddProduct implements the MerchantServiceImpl interface.
func (s *MerchantServiceImpl) AddProduct(ctx context.Context, req *merchant.AddProductReq) (resp *merchant.AddProductResp, err error) {
	resp, err = service.NewAddProductService(ctx).Run(req)

	return resp, err
}

// DeleteProduct implements the MerchantServiceImpl interface.
func (s *MerchantServiceImpl) DeleteProduct(ctx context.Context, req *merchant.DeleteProductReq) (resp *merchant.DeleteProductResp, err error) {
	resp, err = service.NewDeleteProductService(ctx).Run(req)

	return resp, err
}

// UpdateProduct implements the MerchantServiceImpl interface.
func (s *MerchantServiceImpl) UpdateProduct(ctx context.Context, req *merchant.UpdateProductReq) (resp *merchant.UpdateProductResp, err error) {
	resp, err = service.NewUpdateProductService(ctx).Run(req)

	return resp, err
}

// SearchProducts implements the MerchantServiceImpl interface.
func (s *MerchantServiceImpl) SearchProducts(ctx context.Context, req *merchant.SearchProductsReq) (resp *merchant.SearchProductsResp, err error) {
	resp, err = service.NewSearchProductsService(ctx).Run(req)

	return resp, err
}

// ProductDetail implements the MerchantServiceImpl interface.
func (s *MerchantServiceImpl) ProductDetail(ctx context.Context, req *merchant.ProductDetailReq) (resp *merchant.ProductDetailResp, err error) {
	resp, err = service.NewProductDetailService(ctx).Run(req)

	return resp, err
}

// AddMerchant implements the MerchantServiceImpl interface.
func (s *MerchantServiceImpl) AddMerchant(ctx context.Context, req *merchant.AddMerchantReq) (resp *merchant.AddMerchantResp, err error) {
	resp, err = service.NewAddMerchantService(ctx).Run(req)

	return resp, err
}
