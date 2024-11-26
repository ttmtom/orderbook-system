package router

import (
	"github.com/labstack/echo/v4"
	"orderbook/internal/adapter/controller"
	"orderbook/internal/adapter/router/middleware"
)

func InitUserRoute(
	e *echo.Echo,
	midC *middleware.Container,
	uc *controller.UserController,
) {
	user := e.Group("/users")
	{
		user.POST("", uc.Register)
		user.POST("/login", uc.Login)
		authUser := user.Group("/")
		{
			authUser.Use(midC.AuthMiddleware)
			authUser.GET(":idHash", uc.GetUser)
		}
	}
}
