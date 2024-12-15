package repository

import (
	"fmt"
	"gorm.io/gorm"
	"log/slog"
	"orderbook/internal/core/model"
	"orderbook/internal/core/port"
	"orderbook/internal/pkg/security"
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

func (wr *WalletRepository) GetWalletsByUserID(userID uint, filters ...map[string]interface{}) ([]*model.Wallet, error) {
	var wallets []*model.Wallet

	query := wr.db.Model(&model.Wallet{}).
		Where("user_id = ?", userID)
	for _, filter := range filters {
		for key, value := range filter {
			query.Where(fmt.Sprintf("%s = ?", key), value)
		}
	}

	result := query.Find(&wallets)
	if result.Error != nil {
		return nil, result.Error
	}
	return wallets, nil
}

func (wr *WalletRepository) CreateTransaction(transaction *model.Transaction) (*model.Transaction, error) {
	err := wr.db.Transaction(func(tx *gorm.DB) error {
		result := wr.db.Create(&transaction)
		if result.Error != nil {
			return result.Error
		}
		result = wr.db.Model(&transaction).Updates(map[string]interface{}{
			"id_hash": security.HashUserId(transaction.ID),
		})
		if result.Error != nil {
			return result.Error
		}
		newEvent := &model.TransactionEvent{TransactionID: transaction.ID, Type: model.Pending}
		result = wr.db.Create(&newEvent)
		return result.Error
	})

	if err != nil {
		slog.Info("Error on creating transaction", err.Error)
		return nil, err
	}
	return transaction, nil
}
