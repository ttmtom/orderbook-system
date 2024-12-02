package module

import (
	"github.com/go-playground/validator"
	"gorm.io/gorm"
	"orderbook/config"
)

type Container struct {
	UserModule   *UserModule
	CommonModule *CommonModule
}

func InitModuleContainer(
	connection *gorm.DB,
	validator *validator.Validate,
	config *config.Config,
) *Container {
	commonModule := NewCommonModule(connection)
	userModule := NewUserModule(connection, validator, commonModule, config)

	return &Container{
		userModule,
		commonModule,
	}
}
