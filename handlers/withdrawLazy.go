package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/CassioRoos/poc-lazy-loading/models"
	"github.com/CassioRoos/poc-lazy-loading/services"
)

type WithdrawLazy struct {
	service services.WithdrawServiceLazy
}

func NewWithdrawLazy(service services.WithdrawServiceLazy) *WithdrawLazy {
	return &WithdrawLazy{service: service}
}

func (h *WithdrawLazy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	var request models.WithdrawRequest
	err = json.Unmarshal(body, &request)

	if  err :=  h.service.WithdrawLazy(r.Context(),request); err != nil{
		rw.WriteHeader(http.StatusInternalServerError)
	}

	rw.WriteHeader(http.StatusOK)
}

func (h *WithdrawLazy) Method() string{
	return "POST"
}

func (h *WithdrawLazy) Pattern() string{
	return "/withdraw"
}
