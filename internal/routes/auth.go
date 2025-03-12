package routes

import (
	"github.com/ganiyamustafa/bts/internal/controllers"
	"github.com/ganiyamustafa/bts/internal/services"
	"github.com/ganiyamustafa/bts/utils"
	"github.com/labstack/echo/v4"
)

func AuthRoutes(router *echo.Group, handler *utils.Handler) {
	userService := services.UserService{Handler: handler}
	controller := controllers.AuthController{UserService: userService}

	router.POST("/daftar", controller.Register)
	router.POST("/login", controller.Login)
}
