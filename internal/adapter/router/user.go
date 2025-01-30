package router

import (
	"github.com/labstack/echo/v4"
	"orderbook/internal/core/port"
)

func InitUserRouter(
	e *echo.Echo,
	uc port.UserController,
	am port.AuthMiddleware,
) {
	e.POST("/register", uc.Register)
	user := e.Group("/users")
	{
		user.Use(am.HeaderAuthHandler())
		user.GET("", uc.GetMe)
	}
}
