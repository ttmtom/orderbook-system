package controller

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"orderbook/internal/core/port"
	"orderbook/internal/core/service"
	"orderbook/internal/pkg/response"
	"orderbook/pkg/utils"
)

type AuthController struct {
	svc       port.AuthService
	validator *validator.Validate
}

func NewAuthController(validator *validator.Validate, svc port.AuthService) port.AuthController {
	return &AuthController{
		svc:       svc,
		validator: validator,
	}
}

type userLoginRequest struct {
	Email    string `param:"email" validate:"required,email" example:"hi@example.com"`
	Password string `json:"password" validate:"required,cPassword"`
}

func (uc *AuthController) Login(ctx echo.Context) error {
	req, err := utils.ValidateStruct(ctx, uc.validator, new(userLoginRequest))
	if err != nil {
		slog.Info("Validation Error", err.Error())
		return response.FailureResponse(http.StatusBadRequest, err.Error())
	}

	_, tokenSet, err := uc.svc.UserLogin(req.Email, req.Password)
	if err != nil {
		slog.Info("Error during login", err)

		if err.Error() == string(service.Unauthorized) {
			return response.FailureResponse(http.StatusNotFound, err.Error())
		}
		return response.FailureResponse(http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(ctx, http.StatusOK, map[string]string{
		"accessToken":  tokenSet.AccessToken,
		"refreshToken": tokenSet.RefreshToken,
	})
}

type refreshRequest struct {
	Token string `param:"token" validate:"required" example:"hi@example.com"`
}

func (uc *AuthController) RefreshToken(ctx echo.Context) error {
	req, err := utils.ValidateStruct(ctx, uc.validator, new(refreshRequest))
	if err != nil {
		slog.Info("Validation Error", err.Error())
		return response.FailureResponse(http.StatusBadRequest, err.Error())
	}

	token, err := uc.svc.RefreshToken(req.Token)
	if err != nil {
		return response.FailureResponse(http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(ctx, http.StatusOK, map[string]string{
		"accessToken":  token.AccessToken,
		"refreshToken": token.RefreshToken,
	})
}
