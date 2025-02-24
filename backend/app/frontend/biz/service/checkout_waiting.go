package service

import (
	"context"

	"github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/checkout"
	"github.com/tiktokmall/backend/app/frontend/infra/rpc"
	frontendUtils "github.com/tiktokmall/backend/app/frontend/utils"

	rpccheckout "github.com/tiktokmall/backend/rpc_gen/kitex_gen/checkout"
	rpcpayment "github.com/tiktokmall/backend/rpc_gen/kitex_gen/payment"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type CheckoutWaitingService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutWaitingService(Context context.Context, RequestContext *app.RequestContext) *CheckoutWaitingService {
	return &CheckoutWaitingService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutWaitingService) Run(req *checkout.CheckoutReq) (resp map[string]any, err error) {
	// hlog.Info("checkout waiting service run")
	userId := frontendUtils.GetUserIdFromCtx(h.Context)
	_, err = rpc.CheckoutClient.Checkout(h.Context, &rpccheckout.CheckoutReq{
		UserId:        userId,
		Firstname:     req.GetFirstname(),
		Lastname:      req.GetLastname(),
		Email:         req.GetEmail(),
		PaymentMethod: req.GetPayment(),
		Address: &rpccheckout.Address{
			Country:       req.GetCountry(),
			ZipCode:       req.GetZipcode(),
			City:          req.GetCity(),
			State:         req.GetProvince(),
			StreetAddress: req.GetStreet(),
		},
		CreditCard: &rpcpayment.CreditCardInfo{
			CreditCartNumber:          req.GetCardNum(), // 这里写死是因为是card.Validate(true)验证的，不然监测不通过
			CreditCardExpirationYear:  req.GetExpirationYear(),
			CreditCardExpirationMonth: req.GetExpirationMonth(),
			CreditCardCvv:             req.GetCvv(),
		},
	})

	if err != nil {
		hlog.Info("rpc checkout error:", err)
		return nil, err
	}

	return map[string]any{
		"title":    "waiting",
		"redirect": "/checkout/result",
	}, nil
}
