package module

import (
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type Container struct {
	UserModule *UserModule
}

func InitModuleContainer(connection *gorm.DB, validator *validator.Validate) *Container {
	userModule := NewUserModule(connection, validator)

	return &Container{
		userModule,
	}
}
