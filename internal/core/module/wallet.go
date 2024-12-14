package module

import (
	"github.com/go-playground/validator"
	"gorm.io/gorm"
	"orderbook/internal/adapter/database/postgres/repository"
	"orderbook/internal/core/controller"
	"orderbook/internal/core/port"
	"orderbook/internal/core/service"
)

type WalletModule struct {
	Repository port.WalletRepository
	Service    port.WalletService
	Controller port.WalletController
}

func NewWalletModule(
	connection *gorm.DB,
	validator *validator.Validate,
	eventManager port.EventRepository,
	userRepository port.UserRepository,
) *WalletModule {
	wr := repository.NewWalletRepository(connection)
	ws := service.NewWalletService(wr, userRepository)
	wc := controller.NewWalletController(ws, validator)

	eventMap := make(map[string]func(event []byte) error)

	eventMap[string(service.UserRegistrationSuccess)] = ws.OnUserRegistrationSuccess

	eventManager.SetUpGroupConsumer("wallet", eventMap, 500, 5)

	return &WalletModule{
		wr,
		ws,
		wc,
	}
}
