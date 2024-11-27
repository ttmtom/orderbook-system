package router

import (
	"github.com/labstack/echo/v4"
	"orderbook/internal/adapter/controller"
)

func InitUserRoute(
	e *echo.Echo,
	uc *controller.UserController,
	midC *MiddlewareContainer,
) {
	user := e.Group("/users")
	{
		user.POST("", uc.Register)
		user.POST("/login", uc.Login)
		authUser := user.Group("/")
		{
			authUser.Use(midC.AuthMiddleware)
			authUser.GET("", uc.GetUser)
		}
	}
}
