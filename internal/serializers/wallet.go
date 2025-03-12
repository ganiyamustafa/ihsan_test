package serializers

import (
	"github.com/ganiyamustafa/bts/internal/models"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type AddBalanceResponse struct {
	ID       uuid.UUID `json:"id"`
	Balance  float64   `json:"saldo"`
	Currency string    `json:"currency"`
}

func (u AddBalanceResponse) FromModel(m *models.Wallet) *AddBalanceResponse {
	var res AddBalanceResponse
	copier.CopyWithOption(&res, &m, copier.Option{IgnoreEmpty: true})

	return &res
}

type WithdrawBalanceResponse struct {
	ID       uuid.UUID `json:"id"`
	Balance  float64   `json:"saldo"`
	Currency string    `json:"currency"`
}

func (u WithdrawBalanceResponse) FromModel(m *models.Wallet) *WithdrawBalanceResponse {
	var res WithdrawBalanceResponse
	copier.CopyWithOption(&res, &m, copier.Option{IgnoreEmpty: true})

	return &res
}

type GetBalanceResponse struct {
	ID       uuid.UUID `json:"id"`
	Balance  float64   `json:"saldo"`
	Currency string    `json:"currency"`
}

func (u GetBalanceResponse) FromModel(m *models.Wallet) *WithdrawBalanceResponse {
	var res WithdrawBalanceResponse
	copier.CopyWithOption(&res, &m, copier.Option{IgnoreEmpty: true})

	return &res
}
