package router

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"orderbook/config"
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
	middlewareContainer *MiddlewareContainer,
) *Router {
	e := echo.New()
	e.Use(middleware.Logger())
	//e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())

	{
		e.GET("/health", func(e echo.Context) error {
			return e.String(http.StatusOK, "OK")
		})
	}

	InitUserRouter(
		e,
		moduleContainer.UserModule.Controller,
		middlewareContainer,
	)

	return &Router{e, config}
}

func (e *Router) Serve() error {
	return e.echo.Start(fmt.Sprintf("%s:%s", e.config.Host, e.config.Port))
}
