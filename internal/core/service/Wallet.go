package service

import "orderbook/internal/adapter/database/postgres/repository"

type WalletService struct {
	repo *repository.WalletRepository
}

func NewWalletService(repo *repository.WalletRepository) *WalletService {
	return &WalletService{repo: repo}
}

func (ws *WalletService) OnUserRegistrationSuccess() {

}
