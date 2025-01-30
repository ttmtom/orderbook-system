package module

import (
	"github.com/go-playground/validator"
	"gorm.io/gorm"
	"orderbook/config"
	"orderbook/internal/pkg/security"
)

type Container struct {
	EventModule  *EventModule
	UserModule   *UserModule
	CommonModule *CommonModule
	WalletModule *WalletModule
	AuthModule   *AuthModule
	AdminModule  *AdminModule
}

func InitModuleContainer(
	connection *gorm.DB,
	validator *validator.Validate,
	config *config.Config,
) *Container {
	security.InitJwtSecurity(config.AppConfig.SecurityKey)

	eventModule := NewEventModule(config)
	commonModule := NewCommonModule(connection)
	userModule := NewUserModule(connection, validator, eventModule.Repository)
	authModule := NewAuthModule(config.AppConfig, validator, commonModule.Repository, userModule.Repository)
	walletModule := NewWalletModule(connection, validator, eventModule.Repository, userModule.Repository)

	var adminModule *AdminModule = nil
	if config.AppConfig.AdminBuild {
		adminModule = NewAdminModule(config, connection, validator)
	}

	return &Container{
		eventModule,
		userModule,
		commonModule,
		walletModule,
		authModule,
		adminModule,
	}
}
