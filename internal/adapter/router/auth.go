package router

import (
	"github.com/labstack/echo/v4"
	"orderbook/internal/core/port"
)

func InitAuthRouter(
	e *echo.Echo,
	ac port.AuthController,
	am port.AuthMiddleware,
) {
	e.POST("/login", ac.Login)
	e.POST("/refreshToken", ac.RefreshToken, am.HeaderAuthHandler(true))
}
