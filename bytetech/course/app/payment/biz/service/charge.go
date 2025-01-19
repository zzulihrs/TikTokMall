package service

import (
	"bytetech/course/app/payment/biz/dal/mysql"
	"bytetech/course/app/payment/biz/model"
	payment "bytetech/course/rpc_gen/kitex_gen/payment"
	"context"
	"strconv"
	"time"

	"github.com/cloudwego/kitex/pkg/kerrors"
	creditcard "github.com/durango/go-credit-card"
	"github.com/google/uuid"
)

type ChargeService struct {
	ctx context.Context
} // NewChargeService new ChargeService
func NewChargeService(ctx context.Context) *ChargeService {
	return &ChargeService{ctx: ctx}
}

// Run create note info
func (s *ChargeService) Run(req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	// Finish your business logic.
	card := creditcard.Card{
		Number: req.CreditCard.CreditCartNumber,
		Cvv:    strconv.Itoa(int(req.CreditCard.CreditCardCvv)),
		Month:  strconv.Itoa(int(req.CreditCard.CreditCardExpirationMonth)),
		Year:   strconv.Itoa(int(req.CreditCard.CreditCardExpirationYear)),
	}
	err = card.Validate(true)
	if err != nil {
		return nil, kerrors.NewBizStatusError(400, err.Error())
	}
	translationId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	err = model.CreatePaymentLog(mysql.DB, s.ctx, &model.PaymentLog{
		UserId:        req.GetUserId(),
		OrderId:       req.GetOrderId(),
		TransactionId: translationId.String(),
		Amount:        req.GetAmount(),
		PayAt:         time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return &payment.ChargeResp{TransactionId: translationId.String()}, nil
}
