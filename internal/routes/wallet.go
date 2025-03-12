package routes

import (
	"github.com/ganiyamustafa/bts/internal/controllers"
	"github.com/ganiyamustafa/bts/internal/services"
	"github.com/ganiyamustafa/bts/utils"
	"github.com/labstack/echo/v4"
)

func WalletRoutes(router *echo.Group, handler *utils.Handler) {
	userService := services.UserService{Handler: handler}
	walletService := services.WalletService{Handler: handler}
	controller := controllers.WalletController{UserService: userService, WalletService: walletService}

	router.POST("/tabung", controller.AddBalance)
	router.POST("/tarik", controller.WithdrawBalance)

	router.GET("/saldo/:account_number", controller.GetBalance)
}
