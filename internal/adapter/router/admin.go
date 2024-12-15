package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func InitAdminRouter(
	e *echo.Echo,
) {
	admin := e.Group("/admin")
	{
		admin.GET("/", func(e echo.Context) error {
			return e.String(http.StatusOK, "ADMIN")
		})
	}
}
