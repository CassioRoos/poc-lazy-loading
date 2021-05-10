package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/CassioRoos/poc-lazy-loading/models"
	"github.com/CassioRoos/poc-lazy-loading/services"
)

type Withdraw struct {
	service services.WithdrawService
}

func NewWithdraw(service services.WithdrawService) *Withdraw {
	return &Withdraw{service: service}
}

func (h *Withdraw) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	var request models.WithdrawRequest
	err = json.Unmarshal(body, &request)

	if  err :=  h.service.Withdraw(r.Context(), request); err != nil{
		rw.WriteHeader(http.StatusInternalServerError)
	}

	rw.WriteHeader(http.StatusOK)
}

func (h *Withdraw) Method() string{
	return "POST"
}

func (h *Withdraw) Pattern() string{
	return "/withdrawlazy"
}