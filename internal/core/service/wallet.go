package service

import (
	"encoding/json"
	"log/slog"
	"orderbook/internal/core/model"
	"orderbook/internal/core/port"
)

type WalletService struct {
	repo     port.WalletRepository
	userRepo port.UserRepository
}

func NewWalletService(
	repo port.WalletRepository,
	userRepository port.UserRepository,
) port.WalletService {
	return &WalletService{
		repo,
		userRepository,
	}
}

func (ws *WalletService) OnUserRegistrationSuccess(event []byte) error {
	slog.Info("OnUserRegistrationSuccess", event)
	var data UserRegistrationSuccessEvent
	err := json.Unmarshal(event, &data)
	if err != nil {
		slog.Info("failed on unmarshal event", "event", event, "err", err)
		return err
	}

	btcWallet := &model.Wallet{
		UserID:   data.ID,
		Currency: model.BTC,
	}
	btcWallet, err = ws.repo.CreateWallet(btcWallet)

	if err != nil {
		slog.Info("failed on create btc wallet", "data", data, "err", err)
		return err
	}

	ethWallet := &model.Wallet{
		UserID:   data.ID,
		Currency: model.ETH,
	}
	ethWallet, err = ws.repo.CreateWallet(ethWallet)
	if err != nil {
		slog.Info("failed on create eth wallet", "data", data, "err", err)
		return err
	}
	return err
}

func (ws *WalletService) GetWalletsByUserID(userId string) ([]*model.Wallet, error) {
	user, _ := ws.userRepo.GetUserByIdHash(userId)

	wallets, err := ws.repo.GetWalletsByUserID(user.ID)
	if err != nil {
		slog.Info("failed on getting user wallets", "userId", userId, "err", err)
		return nil, err
	}

	return wallets, nil
}

func (ws *WalletService) Deposit(userId string, currency model.CryptoCurrency, source string, amount float64) (*model.Transaction, error) {
	user, _ := ws.userRepo.GetUserByIdHash(userId)

	filters := make(map[string]interface{})
	filters["currency"] = currency

	wallets, err := ws.repo.GetWalletsByUserID(user.ID, filters)
	if err != nil {
		slog.Info("failed on getting user wallets", "userId", userId, "err", err)
		return nil, err
	}

	t := &model.Transaction{
		ToID:        &wallets[0].ID,
		Type:        model.Deposit,
		Amount:      amount,
		Description: source,
	}

	t, err = ws.repo.CreateTransaction(t)

	if err != nil {
		slog.Info("failed on create t", "t", t, "err", err)
		return nil, err
	}

	return t, nil
}

func (ws *WalletService) Withdraw(userId string, currency model.CryptoCurrency, destination string) (*model.Transaction, error) {
	user, _ := ws.userRepo.GetUserByIdHash(userId)

	filters := make(map[string]interface{})
	filters["currency"] = currency

	_, err := ws.repo.GetWalletsByUserID(user.ID, filters)
	return nil, err
}
