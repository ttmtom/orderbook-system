package module

import (
	"github.com/go-playground/validator"
	"gorm.io/gorm"
	"orderbook/internal/adapter/controller"
	"orderbook/internal/adapter/database/postgres/repository"
	"orderbook/internal/adapter/kafka"
	"orderbook/internal/core/service"
)

type UserModule struct {
	Repository   *repository.UserRepository
	Service      *service.UserService
	Controller   *controller.UserController
	KafkaManager *kafka.Manager
}

func NewUserModule(
	connection *gorm.DB,
	validator *validator.Validate,
	commonModule *CommonModule,
	kafkaManager *kafka.Manager,
) *UserModule {
	userRepository := repository.NewUserRepository(connection)
	userService := service.NewUserService(
		userRepository,
		commonModule.Service,
		kafkaManager,
	)
	userController := controller.NewUserController(validator, userService)

	return &UserModule{
		Repository:   userRepository,
		Service:      userService,
		Controller:   userController,
		KafkaManager: kafkaManager,
	}
}
