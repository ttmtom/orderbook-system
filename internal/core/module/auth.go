package module

import (
	"github.com/go-playground/validator"
	"orderbook/config"
	"orderbook/internal/adapter/database/postgres/repository"
	"orderbook/internal/adapter/router/controller"
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
	commonRepo *repository.CommonRepository,
	userRepo *repository.UserRepository,
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
