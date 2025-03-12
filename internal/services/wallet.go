package services

import (
	"net/http"

	"github.com/ganiyamustafa/bts/internal/models"
	"github.com/ganiyamustafa/bts/internal/requests"
	"github.com/ganiyamustafa/bts/utils"
	apperror "github.com/ganiyamustafa/bts/utils/app_error"
	"github.com/jinzhu/copier"
)

type WalletService struct {
	Handler *utils.Handler
}

// create todolist data
func (u *WalletService) CreateWallet(payload *requests.CreateWallet) (*models.Wallet, *apperror.AppError) {
	// copy payload to user models
	var Wallet models.Wallet
	copier.Copy(&Wallet, &payload)

	return &Wallet, apperror.FromError(u.Handler.Postgre.Create(&Wallet).Error)
}

// add balance data
func (u *WalletService) AddBalance(wallet *models.Wallet, payload *requests.AddBalance) (*models.Wallet, *apperror.AppError) {
	*wallet.Balance += payload.Nominal

	return wallet, apperror.FromError(u.Handler.Postgre.Updates(wallet).Error)
}

// add balance data
func (u *WalletService) WithdrawBalance(wallet *models.Wallet, payload *requests.AddBalance) (*models.Wallet, *apperror.AppError) {
	if *wallet.Balance < payload.Nominal {
		return nil, apperror.New("Insufficient Balance").SetHttpCustomStatusCode(http.StatusBadRequest)
	}

	*wallet.Balance -= payload.Nominal
	return wallet, apperror.FromError(u.Handler.Postgre.Updates(wallet).Error)
}
