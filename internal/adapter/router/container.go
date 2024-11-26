package router

import (
	"github.com/labstack/echo/v4"
	"orderbook/internal/adapter/router/middleware"
	"orderbook/internal/core/module"
)

type MiddlewareContainer struct {
	AuthMiddleware func(next echo.HandlerFunc) echo.HandlerFunc
}

func InitMiddlewareContainer(moduleContainer *module.Container) *MiddlewareContainer {
	authMiddleware := middleware.NewAuthMiddleware(
		moduleContainer.CommonModule.Service,
		moduleContainer.UserModule.Service,
	)

	return &MiddlewareContainer{
		AuthMiddleware: authMiddleware.Handler,
	}
}
