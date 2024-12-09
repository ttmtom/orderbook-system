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
	e.Use(am.HeaderAuthHandler())

	e.POST("/deposit", wc.Deposit)
	e.POST("/withdrawal", wc.Withdrawal)
	wallet := e.Group("/wallets")
	{
		wallet.GET("", wc.GetMe)
	}
}
