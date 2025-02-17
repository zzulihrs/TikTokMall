package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	merchant "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/merchant"
	"github.com/tiktokmall/backend/app/frontend/utils"
)

type MerchantGetProductListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewMerchantGetProductListService(Context context.Context, RequestContext *app.RequestContext) *MerchantGetProductListService {
	return &MerchantGetProductListService{RequestContext: RequestContext, Context: Context}
}

func (h *MerchantGetProductListService) Run(req *merchant.MerchantGetProductListReq) (resp utils.H, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}

/*
Response
{
    "code": 200,
    "message": "proident anim id culpa",
    "data": [
        {
            "id": 85,
            "name": "纳喇超",
            "description": "等于或公万整小。设步当采多争。保更们真效务目那。入都今求会活还展。走角再县走有务内少更。",
            "stock": 83,
            "price": "531.38",
            "img_url": "https://loremflickr.com/400/400?lock=6729080248750675"
        },
        {
            "id": 40,
            "name": "寿婷婷",
            "description": "发会写着产件阶场人。情众着声。选给片走。还九值素写增己如。就七照济。物数你区然气族越。边程育公。",
            "stock": 5,
            "price": "801.29",
            "img_url": "https://loremflickr.com/400/400?lock=3451811131915255"
        },
        {
            "id": 98,
            "name": "伏国英",
            "description": "气往温去。火制当角去观每来。",
            "stock": 43,
            "price": "468.14",
            "img_url": "https://loremflickr.com/400/400?lock=8796401005629023"
        }
    ]
}
*/
