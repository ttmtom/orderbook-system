package module

import (
	"gorm.io/gorm"
	"orderbook/internal/adapter/database/postgres/repository"
	"orderbook/internal/adapter/kafka"
	"orderbook/internal/core/service"
)

type WalletModule struct {
	Repository    *repository.WalletRepository
	Service       *service.WalletService
	KafkaConsumer *kafka.ConsumerGroup
}

func NewWalletModule(
	connection *gorm.DB,
	kafkaManager *kafka.Manager,
	userModule *UserModule,
) *WalletModule {
	wr := repository.NewWalletRepository(connection)
	ws := service.NewWalletService(wr)

	eventMap := make(map[string]func(event any))

	eventMap[string(service.UserRegistrationSuccess)] = ws.OnUserRegistrationSuccess

	consumer := kafkaManager.SetUpGroupConsumer("wallet", eventMap)

	return &WalletModule{wr, ws, consumer}
}
