package module

import (
	"github.com/go-playground/validator"
	"gorm.io/gorm"
	"orderbook/config"
	"orderbook/internal/adapter/kafka"
)

type Container struct {
	UserModule   *UserModule
	CommonModule *CommonModule
	WalletModule *WalletModule
	AuthModule   *AuthModule
}

func InitModuleContainer(
	connection *gorm.DB,
	validator *validator.Validate,
	config *config.Config,
	kafkaManager *kafka.Manager,
) *Container {
	commonModule := NewCommonModule(connection)
	userModule := NewUserModule(connection, validator, commonModule, kafkaManager)
	authModule := NewAuthModule(config.AppConfig, validator, commonModule, userModule)
	walletModule := NewWalletModule(connection, kafkaManager, userModule)

	return &Container{
		userModule,
		commonModule,
		walletModule,
		authModule,
	}
}
