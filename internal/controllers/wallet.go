package controllers

import (
	"net/http"

	"github.com/ganiyamustafa/bts/internal/requests"
	"github.com/ganiyamustafa/bts/internal/serializers"
	"github.com/ganiyamustafa/bts/internal/services"
	apperror "github.com/ganiyamustafa/bts/utils/app_error"
	"github.com/labstack/echo/v4"
)

type WalletController struct {
	UserService   services.UserService
	WalletService services.WalletService
}

func (u *WalletController) AddBalance(ctx echo.Context) error {
	var payload requests.AddBalance

	if err := ctx.Bind(&payload); err != nil {
		return ErrorResponse(ctx, apperror.FromError(err).SetHttpCustomStatusCode(http.StatusBadRequest))
	}

	if err := u.UserService.Handler.Validator.Struct(&payload); err != nil {
		return ErrorResponse(ctx, apperror.FromError(err).SetHttpCustomStatusCode(http.StatusBadRequest))
	}

	// validate existing user with wallet
	user, err := u.UserService.GetUserByAccountNumber(payload.AccountNumber)
	if err != nil {
		return ErrorResponse(ctx, apperror.New("Account Number Not Found").SetHttpCustomStatusCode(http.StatusBadRequest))
	}

	// add balance to wallet
	wallet, err := u.WalletService.AddBalance(user.Wallet, &payload)
	if err != nil {
		return ErrorResponse(ctx, err.SetHttpCustomStatusCode(http.StatusBadRequest))
	}

	return SuccessResponse(ctx, serializers.AddBalanceResponse{}.FromModel(wallet), nil, "Add Balance Successfully", http.StatusCreated)
}

func (u *WalletController) WithdrawBalance(ctx echo.Context) error {
	var payload requests.AddBalance

	if err := ctx.Bind(&payload); err != nil {
		return ErrorResponse(ctx, apperror.FromError(err).SetHttpCustomStatusCode(http.StatusBadRequest))
	}

	if err := u.UserService.Handler.Validator.Struct(&payload); err != nil {
		return ErrorResponse(ctx, apperror.FromError(err).SetHttpCustomStatusCode(http.StatusBadRequest))
	}

	// validate existing user with wallet
	user, err := u.UserService.GetUserByAccountNumber(payload.AccountNumber)
	if err != nil {
		return ErrorResponse(ctx, apperror.New("Account Not Found").SetHttpCustomStatusCode(http.StatusBadRequest))
	}

	// withdraw balance to wallet
	wallet, err := u.WalletService.WithdrawBalance(user.Wallet, &payload)
	if err != nil {
		return ErrorResponse(ctx, err.SetHttpCustomStatusCode(http.StatusBadRequest))
	}

	return SuccessResponse(ctx, serializers.WithdrawBalanceResponse{}.FromModel(wallet), nil, "Withdraw Successfully", http.StatusCreated)
}

func (u *WalletController) GetBalance(ctx echo.Context) error {
	accountNumber := ctx.Param("account_number")

	// validate existing user with wallet
	user, err := u.UserService.GetUserByAccountNumber(accountNumber)
	if err != nil {
		return ErrorResponse(ctx, apperror.New("Account Not Found").SetHttpCustomStatusCode(http.StatusBadRequest))
	}

	return SuccessResponse(ctx, serializers.GetBalanceResponse{}.FromModel(user.Wallet), nil, "Withdraw Successfully", http.StatusCreated)
}
