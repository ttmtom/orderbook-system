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
	wallet := e.Group("")
	wallet.Use(am.HeaderAuthHandler())

	wallet.POST("/deposit", wc.Deposit, am.HeaderAuthHandler())
	wallet.POST("/withdrawal", wc.Withdrawal, am.HeaderAuthHandler())
	{
		wallet.GET("/wallets", wc.GetMe)
	}
}
