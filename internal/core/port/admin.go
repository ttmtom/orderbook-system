package port

import (
	"github.com/labstack/echo/v4"
	"orderbook/internal/core/model"
)

type AdminRepository interface {
	GetAdmin(email string) (*model.AdminUser, error)
}

type AdminService interface {
	AdminLogin(email, password string) (*model.AdminUser, string, error)
}

type AdminController interface {
	AdminLogin(ctx echo.Context) error
}

type AdminMiddleware interface {
	AdminAuthHandler() func(next echo.HandlerFunc) echo.HandlerFunc
}
