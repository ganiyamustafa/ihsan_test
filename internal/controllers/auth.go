package controllers

import (
	"net/http"

	"github.com/ganiyamustafa/bts/internal/requests"
	"github.com/ganiyamustafa/bts/internal/serializers"
	"github.com/ganiyamustafa/bts/internal/services"
	"github.com/ganiyamustafa/bts/utils"
	apperror "github.com/ganiyamustafa/bts/utils/app_error"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	UserService   services.UserService
	WalletService services.WalletService
}

func (u *AuthController) Register(ctx echo.Context) error {
	var payload requests.RegisterRequest

	if err := ctx.Bind(&payload); err != nil {
		return ErrorResponse(ctx, apperror.FromError(err).SetHttpCustomStatusCode(http.StatusBadRequest))
	}

	if err := u.UserService.Handler.Validator.Struct(&payload); err != nil {
		return ErrorResponse(ctx, apperror.FromError(err).SetHttpCustomStatusCode(http.StatusBadRequest))
	}

	// validate existing user
	_, err := u.UserService.GetUserByNIKOrPhone(payload.NIK, payload.Phone)
	if err == nil {
		return ErrorResponse(ctx, apperror.New("Nik or Phone has been Used").SetHttpCustomStatusCode(http.StatusBadRequest))
	}

	// create user
	user, err := u.UserService.CreateUser(payload)
	if err != nil {
		return ErrorResponse(ctx, err)
	}

	return SuccessResponse(ctx, serializers.RegisterResponse{AccountNumber: user.AccountNumber}, nil, "Register Successfully", http.StatusOK)
}

func (u *AuthController) Login(ctx echo.Context) error {
	var payload requests.LoginRequest

	if err := ctx.Bind(&payload); err != nil {
		return ErrorResponse(ctx, apperror.FromError(err).SetHttpCustomStatusCode(http.StatusBadRequest))
	}

	if err := u.UserService.Handler.Validator.Struct(&payload); err != nil {
		return ErrorResponse(ctx, apperror.FromError(err).SetHttpCustomStatusCode(http.StatusBadRequest))
	}

	// validate existing user
	user, err := u.UserService.GetUserByAccountNumberNIKAndPhone(payload.AccountNumber, payload.NIK, payload.Phone)
	if err != nil {
		return ErrorResponse(ctx, apperror.New("User Not Found").SetHttpCustomStatusCode(http.StatusBadRequest))
	}

	// generate auth token
	authToken, nErr := utils.EncodeJWT(map[string]string{"id": user.ID.String(), "nik": user.NIK}, []byte(utils.Env("SECRET_KEY")))
	if nErr != nil {
		return ErrorResponse(ctx, apperror.FromError(nErr))
	}

	return SuccessResponse(ctx, serializers.LoginResponse{AuthToken: authToken}, nil, "Register Successfully", http.StatusOK)
}
