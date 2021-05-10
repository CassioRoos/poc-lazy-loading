package handlers

import (
	"net/http"

	v2 "gitlab.com/balance-inc/go-commons/log/v2"
)

type Wisner struct {
}

func newWisner(logger v2.Logger) HandlerOutput {
	return HandlerOutput{
		Handler: &Wisner{

		},
	}
}

func (h *Wisner) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	logger := v2.FromContext(r.Context())
	logger.Info("Info")
	rw.WriteHeader(http.StatusOK)
}

func (h *Wisner) Method() string {
	return "GET"
}

func (h *Wisner) Pattern() string {
	return "/wisner"
}
