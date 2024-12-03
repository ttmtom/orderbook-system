package module

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
	"log/slog"
	"orderbook/config"
	"orderbook/internal/adapter/controller"
	"orderbook/internal/adapter/database/postgres/repository"
	"orderbook/internal/core/service"
)

type UserModule struct {
	Repository *repository.UserRepository
	Service    *service.UserService
	Controller *controller.UserController
	//Consumer   *kafka.Consumer
	Producer *kafka.Producer
}

func NewUserModule(
	connection *gorm.DB,
	validator *validator.Validate,
	commonModule *CommonModule,
	config *config.Config,
) *UserModule {
	userProducer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":        config.KafkaConfig.Brokers,
		"allow.auto.create.topics": true, // @TODO update it for prod
	})
	if err != nil {
		slog.Info("Init User module error", "err", err)
		panic(err)
	}

	userRepository := repository.NewUserRepository(connection)
	userService := service.NewUserService(
		userRepository,
		commonModule.Service,
		userProducer,
	)
	userController := controller.NewUserController(validator, userService)

	return &UserModule{
		Repository: userRepository,
		Service:    userService,
		Controller: userController,
		Producer:   userProducer,
	}
}
