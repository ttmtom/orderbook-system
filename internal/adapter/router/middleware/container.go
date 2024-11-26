package middleware

import (
	"github.com/labstack/echo/v4"
	"orderbook/internal/core/module"
)

type Container struct {
	AuthMiddleware func(next echo.HandlerFunc) echo.HandlerFunc
}

func InitMiddlewareContainer(moduleContainer *module.Container) *Container {
	authMiddleware := NewAuthMiddleware(
		moduleContainer.CommonModule.Service,
		moduleContainer.UserModule.Service,
	)

	return &Container{
		AuthMiddleware: authMiddleware.Handler,
	}
}
