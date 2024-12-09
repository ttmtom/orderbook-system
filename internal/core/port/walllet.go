package port

import (
	"github.com/labstack/echo/v4"
	"orderbook/internal/core/model"
)

type WalletRepository interface {
	CreateWallet(wallet *model.Wallet) (*model.Wallet, error)
}

type WalletService interface {
	OnUserRegistrationSuccess(event []byte) error
}

type WalletController interface {
	Deposit(ctx echo.Context) error
	Withdraw(ctx echo.Context) error
	GetMe(ctx echo.Context) error
}
