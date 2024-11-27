package router

import (
	"github.com/labstack/echo/v4"
	"orderbook/internal/adapter/controller"
)

func InitUserRouter(
	e *echo.Echo,
	uc *controller.UserController,
	midC *MiddlewareContainer,
) {
	e.POST("/login", uc.Login)
	e.POST("/register", uc.Register)

	user := e.Group("/users")
	{
		user.Use(midC.HeaderAuthMiddleware)
		user.GET("", uc.GetUser)
	}
}
