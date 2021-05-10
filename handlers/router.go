package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	v2 "gitlab.com/balance-inc/go-commons/log/v2"
	"gitlab.com/balance-inc/go-commons/middlewares"
	"go.uber.org/fx"
)

var ModuleRouter = fx.Options(
	fx.Provide(
		NewRouter,
	),
)

func NewRouter(params Params, logger  v2.Logger) http.Handler {
	var srvMux = mux.NewRouter()
	mid := middlewares.NewDefaultMiddlewares(logger)
	srvMux.Use(mid.AddDefaultMiddlewaresServices)
	for _, handler := range params.Handlers {
		debugMessage := fmt.Sprintf(
			"Registering %s %s",
			handler.Pattern(),
			handler.Method(),
		)
		logger.Info(debugMessage)

		srvMux.Handle(handler.Pattern(), handler).Methods(handler.Method())
	}
	return srvMux
}
