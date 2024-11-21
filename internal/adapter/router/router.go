package router

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"orderbook/config"
	"orderbook/internal/adapter/handler"
)

type Router struct {
	echo   *echo.Echo
	config *config.HttpConfig
}

type Handler interface {
	Reg(c echo.Context) error
}

func NewRouter(
	config *config.HttpConfig,
	userHandler *handler.UserHandler,
) *Router {
	e := echo.New()

	{
		e.GET("/health", func(e echo.Context) error {
			return e.String(http.StatusOK, "OK")
		})
	}

	user := e.Group("/users")
	{
		user.POST("", userHandler.Register)
		/* TODO add auth
		authUser := user.Group("/").Use(authMiddleware())
		*/
		authUser := user.Group("")
		{
			authUser.GET("/:id", userHandler.GetUser)
		}
	}

	return &Router{e, config}
}

func (e *Router) Serve() error {
	return e.echo.Start(fmt.Sprintf("%s:%s", e.config.Host, e.config.Port))
}
