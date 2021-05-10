package services

import (
	"context"
	"fmt"

	v2 "gitlab.com/balance-inc/go-commons/log/v2"

	"github.com/CassioRoos/poc-lazy-loading/models"
)

type WithdrawServiceLazy struct {
}

func NewWithdrawServiceLazy(logger v2.Logger) WithdrawServiceLazy {
	return WithdrawServiceLazy{}
}

func (sl WithdrawServiceLazy) newServiceFactory(withdrawType models.WithdrawType) models.IWithdraw {
	if !sl.isValidType(withdrawType) {
		panic(fmt.Sprintf("%s: Is not a valid type", withdrawType))
	}

	switch withdrawType {
	case models.WithdrawTypeTabapay:
		return NewWithdrawTabapayService()
	case models.WithdrawTypePayU:
		return NewWithdrawPayUService()
	}
	return nil
}

func (sl WithdrawServiceLazy) isValidType(withdrawType models.WithdrawType) bool {
	for _, t := range models.WithdrawTypes {
		if t == withdrawType {
			return true
		}
	}
	return false
}

func (sl WithdrawServiceLazy) WithdrawLazy(
	ctx context.Context,
	request models.WithdrawRequest) error {
	logger := v2.FromContext(ctx)

	if !sl.isValidType(request.Type) {
		logger.Error(
			"Invalid type for request",
			v2.Values{
				"request.type": request.Type,
			},
		)
		return fmt.Errorf("%s: Is not a valid type", request.Type)
	}

	loggerCtx := v2.WithContext(
		ctx,
		logger,
		v2.Values{
			"request.type": request.Type.String(),
		},
	)

	service := sl.newServiceFactory(request.Type)

	service.Execute(loggerCtx, request)
	return nil
}
