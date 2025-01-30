package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func InitAdminRouter(
	e *echo.Echo,
) {
	admin := e.Group("/admin")
	adminWallet := admin.Group("/wallet")
	{
		adminWallet.GET("/pending", func(e echo.Context) error {
			return e.String(http.StatusOK, "ADMIN")
		})
	}
}
