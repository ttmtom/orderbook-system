package module

import (
	"github.com/go-playground/validator"
	"gorm.io/gorm"
	"orderbook/config"
	"orderbook/internal/adapter/database/postgres/repository"
	"orderbook/internal/core/controller"
	"orderbook/internal/core/middleware"
	"orderbook/internal/core/port"
	"orderbook/internal/core/service"
)

type AdminModule struct {
	Middleware port.AdminMiddleware
	Service    port.AdminService
	Repository port.AdminRepository
	Controller port.AdminController
}

func NewAdminModule(config *config.Config, connection *gorm.DB, validator *validator.Validate) *AdminModule {
	am := middleware.NewAdminMiddleware()
	as := service.NewAdminService()
	ar := repository.NewAdminRepository(connection)
	ac := controller.NewAdminController(validator, as)

	return &AdminModule{
		Middleware: am,
		Service:    as,
		Repository: ar,
		Controller: ac,
	}
}
