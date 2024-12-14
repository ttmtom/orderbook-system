package controller

import (
	"errors"
	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"orderbook/internal/core/model"
	"orderbook/internal/core/port"
	"orderbook/internal/pkg/response"
	"orderbook/internal/pkg/security"
	"orderbook/pkg/utils"
)

type WalletController struct {
	svc       port.WalletService
	validator *validator.Validate
}

func NewWalletController(svc port.WalletService, validator *validator.Validate) *WalletController {
	return &WalletController{
		svc:       svc,
		validator: validator,
	}
}

type depositRequest struct {
	Amount   float64              `json:"amount" validate:"required,gt=0"`
	Currency model.CryptoCurrency `json:"currency" validate:"required,cCryptoCurrencyEnum"`
	Source   string               `json:"source" validate:"required"`
}

func (wc *WalletController) Deposit(ctx echo.Context) error {
	userToken := ctx.Get("user").(*jwt.Token)
	userClaims := userToken.Claims.(*security.UserClaims)

	req, err := utils.ValidateStruct(ctx, wc.validator, new(depositRequest))
	if err != nil {
		slog.Info("Validation Error", err.Error())
		return err
	}

	t, err := wc.svc.Deposit(userClaims.UserID, req.Currency, req.Source, req.Amount)
	if err != nil {
		return err
	}

	return response.SuccessResponse(ctx, http.StatusOK, map[string]interface{}{
		"transaction": t.IDHash,
	})
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
