package module

import (
	"orderbook/internal/core/middleware"
	"orderbook/internal/core/port"
)

type AdminModule struct {
	Middleware port.AdminMiddleware
}

func NewAdminModule() *AdminModule {
	am := middleware.NewAdminMiddleware()
	return &AdminModule{
		Middleware: am,
	}
}
