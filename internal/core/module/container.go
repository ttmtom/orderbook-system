package module

import (
	"github.com/go-playground/validator"
	"gorm.io/gorm"
	"orderbook/config"
)

type Container struct {
	UserModule   *UserModule
	CommonModule *CommonModule
	WalletModule *WalletModule
}

func InitModuleContainer(
	connection *gorm.DB,
	validator *validator.Validate,
	config *config.Config,
) *Container {
	commonModule := NewCommonModule(connection)
	userModule := NewUserModule(connection, validator, commonModule, config)
	walletModule := NewWalletModule(connection)

	return &Container{
		userModule,
		commonModule,
		walletModule,
	}
}
