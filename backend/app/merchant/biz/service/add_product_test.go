package service

import (
	"context"
	"strings"
	"testing"

	"github.com/tiktokmall/backend/app/merchant/biz/util"
	merchant "github.com/tiktokmall/backend/rpc_gen/kitex_gen/merchant"
)

func TestAddProduct_Run(t *testing.T) {
	ctx := context.Background()
	s := NewAddProductService(ctx)
	// init req and assert value

	imgUrl := strings.Join([]string{util.GenerateRandomString(10), util.GenerateRandomString(10)}, "/")
	req := &merchant.AddProductReq{
		MerchantId: 1,
		Product: &merchant.MerchantProductDetailInfo{
			Name:        util.GenerateRandomString(10),
			Description: util.GenerateRandomString(100),
			Stock:       util.GenerateRandomInt32(0, 100),
			Price:       util.GenerateRandomFloat32(0, 100),
			ImgUrl:      imgUrl,
			SliderImgs:  []string{imgUrl},
			Category: []*merchant.Category{
				{
					Id: 1,
				},
				{
					Id: 2,
				},
			},
		},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
