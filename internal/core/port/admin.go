package port

import "github.com/labstack/echo/v4"

type AdminService interface {
}

type AdminController interface {
}

type AdminMiddleware interface {
	AdminAuthHandler() func(next echo.HandlerFunc) echo.HandlerFunc
}
