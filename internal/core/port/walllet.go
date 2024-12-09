package port

import (
	"github.com/labstack/echo/v4"
	"orderbook/internal/core/model"
)

type WalletRepository interface {
	CreateWallet(wallet *model.Wallet) (*model.Wallet, error)
	GetWalletsByUserID(userID uint) ([]*model.Wallet, error)
}

type WalletService interface {
	OnUserRegistrationSuccess(event []byte) error
	GetWalletsByUserID(userID string) ([]*model.Wallet, error)
}

type WalletController interface {
	Deposit(ctx echo.Context) error
	Withdrawal(ctx echo.Context) error
	GetMe(ctx echo.Context) error
}
