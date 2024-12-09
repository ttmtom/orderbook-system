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
	e.POST("/deposit", wc.Deposit, am.HeaderAuthHandler())
	e.POST("/withdrawal", wc.Withdrawal, am.HeaderAuthHandler())
	wallet := e.Group("/wallets")
	{
		wallet.Use(am.HeaderAuthHandler())
		wallet.GET("", wc.GetMe)
	}
}
