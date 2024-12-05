package router

import (
	"github.com/labstack/echo/v4"
	"orderbook/internal/adapter/router/controller"
	"orderbook/internal/core/middleware"
)

func InitUserRouter(
	e *echo.Echo,
	uc *controller.UserController,
	am *middleware.AuthMiddleware,
) {
	e.POST("/register", uc.Register)
	user := e.Group("/users")
	{
		user.Use(am.HeaderAuthHandler())
		user.GET("", uc.GetUser)
	}
}
