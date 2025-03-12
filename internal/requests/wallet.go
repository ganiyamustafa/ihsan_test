package requests

import (
	"github.com/google/uuid"
)

type CreateWallet struct {
	UserID   uuid.UUID `json:"-"`
	Balance  float64   `json:"-"`
	Currency string    `json:"-"`
}

type AddBalance struct {
	AccountNumber string  `json:"no_rekening" validate:"required"`
	Nominal       float64 `json:"nominal" validate:"required"`
}

type WithdrawBalance struct {
	AccountNumber string  `json:"no_rekening" validate:"required"`
	Nominal       float64 `json:"nominal" validate:"required"`
}
