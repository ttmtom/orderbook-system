package module

import (
	"gorm.io/gorm"
	"orderbook/internal/adapter/database/postgres/repository"
)

type WalletModule struct {
	Repository *repository.WalletRepository
}

func NewWalletModule(connection *gorm.DB) *WalletModule {
	wr := repository.NewWalletRepository(connection)

	return &WalletModule{wr}
}
