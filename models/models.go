package models

import (
	"context"
	"encoding/json"
	"fmt"
)

type WithdrawType string

const (
	WithdrawTypeTabapay WithdrawType = "tabapay"
	WithdrawTypePayU    WithdrawType = "payu"
)

var WithdrawTypes = [...]WithdrawType{
	WithdrawTypeTabapay,
	WithdrawTypePayU,
}

func (w WithdrawType) String() string {
	return string(w)
}

func (w WithdrawType) GetType(value string) (WithdrawType, error) {
	for _, withdrawType := range WithdrawTypes {
		if withdrawType.String() == value {
			return withdrawType, nil
		}
	}
	return WithdrawType(""), fmt.Errorf("%s: Is not a withdraw type valid", value)
}

type WithdrawRequest struct {
	EmployerID string       `json:"employer_id"`
	Type       WithdrawType `json:"type"`

	// Check payload value
	Payload json.RawMessage `json:"payload"`
}

type WithdrawPayU struct {
	Amount    int    `json:"amount"`
	Operation string `json:"operation"`
}

type WithdrawTabapay struct {
	Currency  string  `json:"currency"`
	Amount    float64 `json:"amount"`
	Sender    string  `json:"sender"`
	Recipient string  `json:"recipient"`
}

type IWithdraw interface {
	Execute(ctx context.Context, request WithdrawRequest) error
	GetType() WithdrawType
}
