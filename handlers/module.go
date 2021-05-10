package handlers

import (
	v2 "gitlab.com/balance-inc/go-commons/log/v2"
	"go.uber.org/fx"

	"github.com/CassioRoos/poc-lazy-loading/services"
)

var Module = fx.Options(
	ModuleRouter,
	fx.Provide(
		newWithdraw,
		newWithdrawLazy,
		newWisner,
	),
)

func newWithdraw() HandlerOutput {
	withdrawService := services.NewWithdrawService()
	return HandlerOutput{
		Handler: NewWithdraw(withdrawService),
	}
}

func newWithdrawLazy(logger v2.Logger) HandlerOutput {
	withdrawServicelazy := services.NewWithdrawServiceLazy(logger)
	return HandlerOutput{
		Handler: NewWithdrawLazy(withdrawServicelazy),
	}

}
