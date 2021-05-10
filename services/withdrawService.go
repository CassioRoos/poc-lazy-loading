package services

import (
	"context"
	"fmt"

	v2 "gitlab.com/balance-inc/go-commons/log/v2"

	"github.com/CassioRoos/poc-lazy-loading/models"
)

type WithdrawService struct {
	Services map[models.WithdrawType]models.IWithdraw
}

func NewWithdrawService() WithdrawService {
	services := make(map[models.WithdrawType]models.IWithdraw)
	tabapay := NewWithdrawTabapayService()
	payU := NewWithdrawPayUService()
	services[tabapay.GetType()] = tabapay
	services[payU.GetType()] = payU

	return WithdrawService{Services: services}
}

func (w WithdrawService) Withdraw(ctx context.Context, request models.WithdrawRequest) error {
	logger := v2.FromContext(ctx)

	service, ok := w.Services[request.Type]
	if !ok {
		return fmt.Errorf("%s: Is not a valid type", request.Type)
	}

	loggerCtx := v2.WithContext(
		ctx,
		logger,
		v2.Values{
			"request.type": request.Type.String(),
		},
	)

	service.Execute(loggerCtx, request)
	return nil
}
