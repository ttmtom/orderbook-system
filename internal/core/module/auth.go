package module

import (
	"github.com/go-playground/validator"
	"orderbook/config"
	"orderbook/internal/adapter/router/controller"
	"orderbook/internal/core/middleware"
	"orderbook/internal/core/port"
	"orderbook/internal/core/service"
)

type AuthModule struct {
	Controller port.AuthController
	Service    port.AuthService
	Middleware port.AuthMiddleware
}

func NewAuthModule(
	config *config.AppConfig,
	validator *validator.Validate,
	commonRepo port.CommonRepository,
	userRepo port.UserRepository,
) *AuthModule {
	as := service.NewAuthService(commonRepo, userRepo)
	ac := controller.NewAuthController(validator, as)
	mid := middleware.NewAuthMiddleware(
		config,
		//commonModule.Repository,
		as,
	)

	return &AuthModule{ac, as, mid}
}
