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
	appconfig *config.AppConfig,
	moduleContainer *module.Container,
) *Router {
	e := echo.New()
	e.Use(middleware.Logger())
	//e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())

	{
		e.GET("/_healthcheck", func(e echo.Context) error {
			return e.String(http.StatusOK, "OK")
		})
	}

	InitUserRouter(
		e,
		moduleContainer.UserModule.Controller,
		moduleContainer.AuthModule.Middleware,
	)
	InitAuthRouter(
		e,
		moduleContainer.AuthModule.Controller,
		moduleContainer.AuthModule.Middleware,
	)
	InitWalletRouter(
		e,
		moduleContainer.WalletModule.Controller,
		moduleContainer.AuthModule.Middleware,
	)

	if appconfig.AdminBuild {
		InitAdminRouter(e)
	}

	return &Router{e, config}
}

func (e *Router) Serve() error {
	return e.echo.Start(fmt.Sprintf("%s:%s", e.config.Host, e.config.Port))
}
