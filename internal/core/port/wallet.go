package port

import (
	"github.com/labstack/echo/v4"
	"orderbook/internal/core/model"
)

type WalletRepository interface {
	CreateWallet(wallet *model.Wallet) (*model.Wallet, error)
	CreateTransaction(transaction *model.Transaction) (*model.Transaction, error)
	GetWalletsByUserID(userID uint, filters ...map[string]interface{}) ([]*model.Wallet, error)
}

type WalletService interface {
	OnUserRegistrationSuccess(event []byte) error
	GetWalletsByUserID(userID string) ([]*model.Wallet, error)
	Deposit(userId string, currency model.CryptoCurrency, source string, amount float64) (*model.Transaction, error)
}

type WalletController interface {
	Deposit(ctx echo.Context) error
	Withdrawal(ctx echo.Context) error
	GetMe(ctx echo.Context) error
}
