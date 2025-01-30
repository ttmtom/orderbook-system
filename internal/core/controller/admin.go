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

type AdminController struct {
	validator *validator.Validate
	svc       port.AdminService
}

func NewAdminController(
	validator *validator.Validate,
	svc port.AdminService,
) port.AdminController {
	return &AdminController{
		validator,
		svc,
	}
}

type adminLoginRequest struct {
	Email    string `param:"email" validate:"required,email" example:"hi@example.com"`
	Password string `param:"password" validate:"required,email" example:"hi@example.com"`
}

func (ac AdminController) AdminLogin(ctx echo.Context) error {
	req, err := utils.ValidateStruct(ctx, ac.validator, new(adminLoginRequest))
	if err != nil {
		slog.Info("Validation Error", err.Error())
		return response.FailureResponse(http.StatusBadRequest, err.Error())
	}

	_, token, err := ac.svc.AdminLogin(req.Email, req.Password)
	if err != nil {
		slog.Info("Error during login", err)

		if err.Error() == string(service.Unauthorized) {
			return response.FailureResponse(http.StatusNotFound, err.Error())
		}
		return response.FailureResponse(http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(ctx, http.StatusOK, map[string]string{
		"accessToken": token,
	})
}
