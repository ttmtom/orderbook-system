package module

import (
	"gorm.io/gorm"
	"orderbook/internal/adapter/database/postgres/repository"
	"orderbook/internal/core/port"
	"orderbook/internal/core/service"
)

type WalletModule struct {
	Repository port.WalletRepository
	Service    port.WalletService
}

func NewWalletModule(
	connection *gorm.DB,
	eventManager port.EventRepository,
) *WalletModule {
	wr := repository.NewWalletRepository(connection)
	ws := service.NewWalletService(wr)

	eventMap := make(map[string]func(event []byte) error)

	eventMap[string(service.UserRegistrationSuccess)] = ws.OnUserRegistrationSuccess

	eventManager.SetUpGroupConsumer("wallet", eventMap, 500, 5)

	return &WalletModule{
		wr,
		ws,
	}
}
