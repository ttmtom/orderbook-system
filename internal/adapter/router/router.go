package router

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"orderbook/config"
	"orderbook/internal/adapter/router/middleware"
	"orderbook/internal/core/module"
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
	moduleContainer *module.Container,
	middlewareContainer *middleware.Container,
) *Router {
	e := echo.New()
	{
		e.GET("/health", func(e echo.Context) error {
			return e.String(http.StatusOK, "OK")
		})
	}

	InitUserRoute(e, middlewareContainer, moduleContainer.UserModule.Controller)

	return &Router{e, config}
}

func (e *Router) Serve() error {
	return e.echo.Start(fmt.Sprintf("%s:%s", e.config.Host, e.config.Port))
}
