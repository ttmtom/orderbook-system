package router

import (
	"github.com/labstack/echo/v4"
	"orderbook/internal/core/port"
)

func InitWalletRouter(
	e *echo.Echo,
	wc port.WalletController,
	am port.AuthMiddleware,
) {
	wallet := e.Group("/wallet")
	wallet.Use(am.HeaderAuthHandler())

	wallet.GET("", wc.GetMe)
	wallet.POST("/deposit", wc.Deposit)
	wallet.POST("/withdrawal", wc.Withdrawal)
}
