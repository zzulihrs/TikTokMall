package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	merchant "github.com/tiktokmall/backend/rpc_gen/kitex_gen/merchant"
)

func TestAddMerchant_Run(t *testing.T) {
	ctx := context.Background()
	s := NewAddMerchantService(ctx)
	// init req and assert value

	// todo: edit your unit test
	testCases := []struct {
		caseName string
		args     merchant.AddMerchantReq
	}{
		{
			caseName: "test",
			args: merchant.AddMerchantReq{
				Uid:      100,
				Username: "test",
				Email:    "EMA",
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.caseName, func(t *testing.T) {
			resp, err := s.Run(&tc.args)
			t.Logf("err: %v", err)
			t.Logf("resp: %v", resp)
			// 若 uid 或者 Email重复，则报错
			require.Error(t, err)
		})
	}

}
