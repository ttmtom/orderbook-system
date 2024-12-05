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
	userModule := NewUserModule(connection, validator, kafkaManager)
	authModule := NewAuthModule(config.AppConfig, validator, commonModule.Repository, userModule.Repository)
	walletModule := NewWalletModule(connection, kafkaManager)

	return &Container{
		userModule,
		commonModule,
		walletModule,
		authModule,
	}
}
