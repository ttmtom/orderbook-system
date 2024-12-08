package port

import "orderbook/internal/core/model"

type WalletRepository interface {
	CreateWallet(wallet *model.Wallet) (*model.Wallet, error)
}

type WalletService interface {
	OnUserRegistrationSuccess(event []byte) error
}
