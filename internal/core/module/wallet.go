package module

import (
	"gorm.io/gorm"
	"orderbook/internal/adapter/database/postgres/repository"
	"orderbook/internal/adapter/kafka"
	"orderbook/internal/core/service"
)

type WalletModule struct {
	Repository *repository.WalletRepository
	Service    *service.WalletService
}

func NewWalletModule(connection *gorm.DB, kafkaManager *kafka.Manager) *WalletModule {
	wr := repository.NewWalletRepository(connection)
	ws := service.NewWalletService(wr)

	return &WalletModule{wr, ws}
}
