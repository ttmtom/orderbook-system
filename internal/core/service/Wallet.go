package service

import (
	"encoding/json"
	"log/slog"
	"orderbook/internal/adapter/database/postgres/repository"
	"orderbook/internal/core/model"
)

type WalletService struct {
	repo        *repository.WalletRepository
	userService *UserService
}

func NewWalletService(
	repo *repository.WalletRepository,
	userService *UserService,
) *WalletService {
	return &WalletService{
		repo,
		userService,
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

	user, err := ws.userService.GetUserById(data.ID)
	if err != nil {
		slog.Info("failed on get user", "data", data, "err", err)
		return err
	}

	btcWallet := &model.Wallet{
		User:     *user,
		Currency: model.BTC,
	}
	btcWallet, err = ws.repo.CreateWallet(btcWallet)

	if err != nil {
		slog.Info("failed on create btc wallet", "data", data, "err", err)
		return err
	}

	ethWallet := &model.Wallet{
		User:     *user,
		Currency: model.ETH,
	}
	ethWallet, err = ws.repo.CreateWallet(ethWallet)
	if err != nil {
		slog.Info("failed on create eth wallet", "data", data, "err", err)
		return err
	}
	return err
}
