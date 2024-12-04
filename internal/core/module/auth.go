package module

import (
	"github.com/go-playground/validator"
	"orderbook/config"
	"orderbook/internal/adapter/controller"
	"orderbook/internal/core/middleware"
	"orderbook/internal/core/service"
)

type AuthModule struct {
	Controller *controller.AuthController
	Service    *service.AuthService
	Middleware *middleware.AuthMiddleware
}

func NewAuthModule(
	config *config.AppConfig,
	validator *validator.Validate,
	commonModule *CommonModule,
	userModule *UserModule,
) *AuthModule {
	as := service.NewAuthService(commonModule.Service, userModule.Service, userModule.Repository)
	ac := controller.NewAuthController(validator, as)
	mid := middleware.NewAuthMiddleware(
		config,
		commonModule.Service,
		as,
	)

	return &AuthModule{ac, as, mid}
}
