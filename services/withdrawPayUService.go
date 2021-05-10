package services

import (
	"context"
	"encoding/json"

	v2 "gitlab.com/balance-inc/go-commons/log/v2"

	"github.com/CassioRoos/poc-lazy-loading/models"
)

type WithdrawPayUService struct {
}

func NewWithdrawPayUService() WithdrawPayUService {
	return WithdrawPayUService{}
}

func (w WithdrawPayUService) Execute(
	ctx context.Context,
	request models.WithdrawRequest,
) error {
	logger := v2.FromContext(ctx)
	var payload models.WithdrawPayU
	err := json.Unmarshal(request.Payload, &payload)
	if err != nil {
		logger.Error(
			"Error umarsheling payload",
			v2.Values{"error": err},
		)
		return err
	}

	logger.Info("PayU service", v2.Values{"payload": payload})
	return nil
}

func (w WithdrawPayUService) GetType() models.WithdrawType {
	return models.WithdrawTypePayU
}
