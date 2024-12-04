package repository

import (
	"gorm.io/gorm"
	"orderbook/internal/core/model"
)

type WalletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {
	return &WalletRepository{db: db}
}

func (wr *WalletRepository) CreateWallet(wallet *model.Wallet) (*model.Wallet, error) {
	result := wr.db.Create(&wallet)
	if result.Error != nil {
		return nil, result.Error
	}
	return wallet, nil
}
