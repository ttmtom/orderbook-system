package controller

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"orderbook/internal/core/port"
	"orderbook/internal/pkg/response"
	"orderbook/internal/pkg/security"
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

type DepositRequest struct {
	Amount   float64 `json:"amount" binding:"required"`
	Currency string  `json:"currency" binding:"required"`
	Source   string  `json:"source" binding:"required"`
}

func (wc *WalletController) Deposit(ctx echo.Context) error {
	//userToken := ctx.Get("user").(*jwt.Token)
	//userClaims := userToken.Claims.(*security.UserClaims)

	return errors.New("TODO")
}

type WithdrawRequest struct {
	Amount     float64 `json:"amount" binding:"required"`
	Currency   string  `json:"currency" binding:"required"`
	Departures string  `json:"departures" binding:"required"`
}

func (wc *WalletController) Withdrawal(ctx echo.Context) error {
	//userToken := ctx.Get("user").(*jwt.Token)
	//userClaims := userToken.Claims.(*security.UserClaims)

	return errors.New("TODO")
}

func (wc *WalletController) GetMe(ctx echo.Context) error {
	userToken := ctx.Get("user").(*jwt.Token)
	userClaims := userToken.Claims.(*security.UserClaims)

	wallets, err := wc.svc.GetWalletsByUserID(userClaims.UserID)

	if err != nil {
		slog.Info("Error on get user wallets", err)
		return response.FailureResponse(http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(ctx, http.StatusOK, wallets)
}
