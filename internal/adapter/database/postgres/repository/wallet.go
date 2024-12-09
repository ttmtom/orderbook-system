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

func (wr *WalletRepository) GetWalletsByUserID(userID uint) ([]*model.Wallet, error) {
	var wallets []*model.Wallet

	result := wr.db.Model(&model.Wallet{}).
		Where("user_id = ?", userID).
		Find(&wallets)
	if result.Error != nil {
		return nil, result.Error
	}
	return wallets, nil
}
