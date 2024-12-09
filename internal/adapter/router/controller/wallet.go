package controller

import (
	"errors"
	"github.com/labstack/echo/v4"
	"orderbook/internal/core/port"
)

type WalletController struct {
	svc  port.WalletService
	repo port.WalletRepository
}

func NewWalletController(svc port.WalletService, repo port.WalletRepository) *WalletController {
	return &WalletController{
		svc:  svc,
		repo: repo,
	}
}

func (wc *WalletController) Deposit(ctx echo.Context) error {
	return errors.New("TODO")
}

func (wc *WalletController) Withdraw(ctx echo.Context) error {
	return errors.New("TODO")
}

func (wc *WalletController) GetMe(ctx echo.Context) error {
	return errors.New("TODO")
}
