package module

import (
	"github.com/go-playground/validator"
	"gorm.io/gorm"
	"orderbook/internal/adapter/controller"
	"orderbook/internal/adapter/database/postgres/repository"
	"orderbook/internal/core/service"
)

type UserModule struct {
	Repository *repository.UserRepository
	Service    *service.UserService
	Controller *controller.UserController
}

func NewUserModule(
	connection *gorm.DB,
	validator *validator.Validate,
	commonModule *CommonModule,
) *UserModule {
	userRepository := repository.NewUserRepository(connection)
	userService := service.NewUserService(userRepository, commonModule.Service)
	userController := controller.NewUserController(validator, userService)

	return &UserModule{
		Repository: userRepository,
		Service:    userService,
		Controller: userController,
	}
}
