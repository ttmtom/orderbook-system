package router

import (
	"github.com/labstack/echo/v4"
	"orderbook/config"
	"orderbook/internal/adapter/router/middleware"
	"orderbook/internal/core/module"
)

type MiddlewareContainer struct {
	CookiesAuthMiddleware func(next echo.HandlerFunc) echo.HandlerFunc
	HeaderAuthMiddleware  func(next echo.HandlerFunc) echo.HandlerFunc
}

func InitMiddlewareContainer(config *config.AppConfig, moduleContainer *module.Container) *MiddlewareContainer {
	authMiddleware := middleware.NewAuthMiddleware(
		config,
		moduleContainer.CommonModule.Service,
		moduleContainer.UserModule.Service,
	)

	return &MiddlewareContainer{
		CookiesAuthMiddleware: authMiddleware.CookiesAuthHandler,
		HeaderAuthMiddleware:  authMiddleware.HeaderAuthHandler(),
	}
}
