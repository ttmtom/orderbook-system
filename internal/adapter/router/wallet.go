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
	e.POST("/deposit", wc.Deposit)
	e.POST("/withdraw", wc.Withdraw)
	wallet := e.Group("/wallet")
	{
		wallet.Use(am.HeaderAuthHandler())
		wallet.GET("", wc.GetMe)
	}
}
