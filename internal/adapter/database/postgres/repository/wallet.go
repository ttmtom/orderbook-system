package repository

import (
	"gorm.io/gorm"
	"orderbook/internal/core/model"
	"orderbook/internal/core/port"
)

type WalletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) port.WalletRepository {
	return &WalletRepository{db: db}
}

func (wr *WalletRepository) CreateWallet(wallet *model.Wallet) (*model.Wallet, error) {
	result := wr.db.Create(&wallet)
	if result.Error != nil {
		return nil, result.Error
	}
	return wallet, nil
}
