package service

import (
	"context"
	"fmt"

	"log"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	"github.com/tiktokmall/backend/app/frontend/biz/utils"
	merchant "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/merchant"
	"github.com/tiktokmall/backend/app/frontend/infra/rpc"
	frontendUtils "github.com/tiktokmall/backend/app/frontend/utils"

	rpcmerchant "github.com/tiktokmall/backend/rpc_gen/kitex_gen/merchant"
)

type MerchantAuthService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewMerchantAuthService(Context context.Context, RequestContext *app.RequestContext) *MerchantAuthService {
	return &MerchantAuthService{RequestContext: RequestContext, Context: Context}
}

/*
 */

func (h *MerchantAuthService) Run(req *merchant.MerchantAuthReq) (resp utils.H, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()

	uid := frontendUtils.GetUserIdFromCtx(h.Context)
	email := frontendUtils.GetEmailFromCtx(h.Context)
	if uid == 0 {
		return nil, fmt.Errorf("未登录")
	}
	// 1. 调用 merchantrpc 获取店家信息
	merchantResp, err := rpc.MerchantClient.GetMerchant(h.Context, &rpcmerchant.GetMerchantReq{
		Id: int64(uid),
	})
	log.Printf("merchantResp: %+v", merchantResp)
	if err != nil {
		return nil, err
	}
	if email != merchantResp.Email {
		return nil, fmt.Errorf("认证失败，email 不匹配")
	}

	// 2. 定义新的 session
	// 获取 请求的上下文 得到 session
	session := sessions.Default(h.RequestContext)
	session.Clear() // 清空 session
	// 利用 redis 存储 user_id
	// 设置 Cookie
	// Cookie
	session.Set("user_id", int32(uid))
	session.Set("username", merchantResp.Username)
	session.Set("email", email)
	session.Set("merchant_id", merchantResp.Id)
	log.Printf("session: set merchant_id: %v", merchantResp.Id)
	err = session.Save()
	if err != nil {
		return nil, err
	}
	// 2. 包装店家信息生成 token
	//token := utils.GenerateJWT(merchantResp.Id, merchantResp.Username)
	//// 3. 返回 token 和店家信息 response
	//resp = utils.H{
	//	"code":    200,
	//	"message": "OK",
	//	"token":   token,
	//	"merchant_info": utils.H{
	//		"id":       merchantResp.Id,
	//		"username": merchantResp.Username,
	//		"email":    merchantResp.Email,
	//		"userId":   frontendUtils.GetUserIdFromCtx(h.Context),
	//	},
	//}
	return
}

/*
Response
{
    "code": 61,
    "message": "nulla reprehenderit proident",
    "token": "aute consectetur Ut",
    "merchant_info": {
        "id": 98,
        "username": "戊国良",
        "email": "nlrlqi.px58@gmail.com"
    }
}
*/
