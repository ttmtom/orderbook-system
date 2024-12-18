package middleware

import (
	"github.com/labstack/echo/v4"
	"orderbook/internal/core/port"
)

type AdminMiddleware struct {
}

func (a *AdminMiddleware) AdminAuthHandler() func(next echo.HandlerFunc) echo.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func NewAdminMiddleware() port.AdminMiddleware {
	return &AdminMiddleware{}
}
