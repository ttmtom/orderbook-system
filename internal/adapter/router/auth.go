package router

import (
	"github.com/labstack/echo/v4"
	"orderbook/internal/adapter/controller"
)

func InitAuthRouter(
	e *echo.Echo,
	ac *controller.AuthController,
) {
	e.POST("/login", ac.Login)
	e.POST("/refreshToken", ac.RefreshToken)
}
