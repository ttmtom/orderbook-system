package handler

import (
	"github.com/go-playground/validator"
	"orderbook/internal/pkg/response"

	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"orderbook/internal/core/service"
)

type UserHandler struct {
	svc       *service.UserService
	validator *validator.Validate
}

func NewUserHandler(validator *validator.Validate, svc *service.UserService) *UserHandler {
	return &UserHandler{
		svc,
		validator,
	}
}

type registerRequest struct {
	Email    string `json:"email" validate:"required,email" example:"test@example.com"`
	Password string `json:"password" validate:"required,cPassword"`
}

func (uh *UserHandler) Register(ctx echo.Context) error {
	req := new(registerRequest)
	if err := ctx.Bind(req); err != nil {
		slog.Info("Http error", err.Error())
		return response.FailureResponse(http.StatusBadRequest, err.Error())
	}
	if err := uh.validator.Struct(req); err != nil {
		slog.Error("Validation error", err)
		return response.FailureResponse(http.StatusBadRequest, echo.Map{
			"Message": "Invalid input",
			"Error":   err.Error(),
		})
	}

	user, err := uh.svc.UserRegistration(req.Email, req.Password)

	if err != nil {
		slog.Info("Error during registration", err)
		if err.Error() == string(service.EmailAlreadyExist) {
			return response.FailureResponse(http.StatusConflict, err.Error())
		}
		return response.FailureResponse(http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(ctx, http.StatusOK, user)
}

type getUserRequest struct {
	ID uint `param:"id" validate:"required,min=1" example:"1"`
}

func (uh *UserHandler) GetUser(ctx echo.Context) error {
	var req getUserRequest
	if err := ctx.Bind(&req); err != nil {
		slog.Info("Invalid request body", err)
		return err
	}

	user, err := uh.svc.GetUserInformation(req.ID)
	if err != nil {
		slog.Info("Error during registration", err)

		if err.Error() == string(service.UserNotFound) {
			return response.FailureResponse(http.StatusNotFound, err.Error())
		}
		return response.FailureResponse(http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(ctx, http.StatusOK, user)
}
