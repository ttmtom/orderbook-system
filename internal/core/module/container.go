package module

import (
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type Container struct {
	UserModule   *UserModule
	CommonModule *CommonModule
}

func InitModuleContainer(connection *gorm.DB, validator *validator.Validate) *Container {
	commonModule := NewCommonModule(connection)
	userModule := NewUserModule(connection, validator, commonModule)

	return &Container{
		userModule,
		commonModule,
	}
}
