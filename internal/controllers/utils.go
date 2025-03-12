package controllers

import (
	"net/http"
	"time"

	"github.com/ganiyamustafa/bts/internal/serializers"
	apperror "github.com/ganiyamustafa/bts/utils/app_error"
	"github.com/labstack/echo/v4"
)

func ErrorResponse(ctx echo.Context, err *apperror.AppError) error {
	errMessage := err.Error()

	if err.HttpStatusCode() == http.StatusInternalServerError {
		errMessage = "something went wrong with our server"
	}

	return ctx.JSON(err.HttpStatusCode(), serializers.BaseResponse{
		Status:     err.HttpStatusMessage(),
		StatusCode: err.HttpStatusCode(),
		Message:    errMessage,
		Timestamp:  time.Now(),
	})
}

func SuccessResponse(ctx echo.Context, data any, meta *serializers.MetaResponse, message string, statusCode int) error {
	return ctx.JSON(statusCode, serializers.BaseResponse{
		Status:     "success",
		StatusCode: statusCode,
		Message:    message,
		Timestamp:  time.Now(),
		Data:       data,
		Meta:       meta,
	})
}

func TransactionFunc() {

}
