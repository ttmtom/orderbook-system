package module

import (
	"github.com/go-playground/validator"
	"gorm.io/gorm"
	"orderbook/internal/adapter/database/postgres/repository"
	"orderbook/internal/adapter/router/controller"
	"orderbook/internal/core/port"
	"orderbook/internal/core/service"
)

type UserModule struct {
	Repository port.UserRepository
	Service    port.UserService
	Controller port.UserController
}

func NewUserModule(
	connection *gorm.DB,
	validator *validator.Validate,
	eventRepository port.EventRepository,
) *UserModule {
	userRepository := repository.NewUserRepository(connection)
	userService := service.NewUserService(
		userRepository,
		eventRepository,
	)
	userController := controller.NewUserController(validator, userService)

	return &UserModule{
		Repository: userRepository,
		Service:    userService,
		Controller: userController,
	}
}
