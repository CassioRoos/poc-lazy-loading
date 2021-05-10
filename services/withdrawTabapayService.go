package services

import (
	"context"
	"encoding/json"

	v2 "gitlab.com/balance-inc/go-commons/log/v2"

	"github.com/CassioRoos/poc-lazy-loading/models"
)

type WithdrawTabapayService struct {
}

func NewWithdrawTabapayService() WithdrawTabapayService {
	return WithdrawTabapayService{}
}

func (w WithdrawTabapayService) Execute(
	ctx context.Context,
	request models.WithdrawRequest,
) error {
	logger := v2.FromContext(ctx)
	var payload models.WithdrawTabapay
	err := json.Unmarshal(request.Payload, &payload)
	if err != nil {
		logger.Error(
			"Error umarsheling payload",
			v2.Values{"error": err},
		)
		return err
	}

	logger.Info("Tabapay service", v2.Values{"payload": payload})
	return nil
}

func (w WithdrawTabapayService) GetType() models.WithdrawType {
	return models.WithdrawTypeTabapay
}
