package service

import (
	"context"
	"github.com/tiktokmall/backend/app/frontend/biz/utils"
	rpcmerchant "github.com/tiktokmall/backend/rpc_gen/kitex_gen/merchant"

	auth "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/auth"
	"github.com/tiktokmall/backend/app/frontend/infra/rpc"
	"github.com/tiktokmall/backend/rpc_gen/kitex_gen/user"

	"github.com/cloudwego/hertz/pkg/app"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	rpc_resp, err := rpc.UserClient.Login(h.Context, &user.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})
	// 没有登陆成功
	if err != nil {
		return nil, err
	}
	merchantResp, err := rpc.MerchantClient.GetMerchant(h.Context, &rpcmerchant.GetMerchantReq{
		Id: rpc_resp.UserId,
	})
	merchantId := int64(0)
	if err == nil {
		merchantId = merchantResp.Id
	}
	tokenString, err := utils.GenerateJWT(rpc_resp.UserId, rpc_resp.Username, rpc_resp.Email, merchantId)
	if err == nil {
		return map[string]any{
			"status":  200,
			"token":   tokenString,
			"message": "OK1",
		}, err
	} else {
		return map[string]any{
			"status":  500,
			"message": err,
		}, err
	}

}
