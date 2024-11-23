package handler

import (
	"github.com/go-playground/validator"
	"orderbook/internal/core/model"
	"orderbook/internal/pkg/response"
	"orderbook/pkg/utils"
	"time"

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

type useResponse struct {
	ID          string    `json:"id"`
	Email       string    `json:"email"`
	DisplayName *string   `json:"displayName"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (uh *UserHandler) formatUserResponse(user *model.User) *useResponse {
	return &useResponse{
		Email:       user.Email,
		ID:          user.IDHash,
		DisplayName: user.DisplayName,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
}

type registerRequest struct {
	Email    string `json:"email" validate:"required,email" example:"test@example.com"`
	Password string `json:"password" validate:"required,cPassword"`
}

func (uh *UserHandler) Register(ctx echo.Context) error {
	req, err := utils.ValidateStruct(ctx, uh.validator, new(registerRequest))
	if err != nil {
		slog.Info("Validation Error", err.Error())
		return err
	}

	user, err := uh.svc.UserRegistration(req.Email, req.Password)
	if err != nil {
		slog.Info("Error during registration", err)
		if err.Error() == string(service.EmailAlreadyExist) {
			return response.FailureResponse(http.StatusConflict, err.Error())
		}
		return response.FailureResponse(http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(ctx, http.StatusOK, uh.formatUserResponse(user))
}

type getUserRequest struct {
	ID string `param:"id" validate:"required,min=1" example:"1"`
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
	return response.SuccessResponse(ctx, http.StatusOK, uh.formatUserResponse(user))
}

type userLoginRequest struct {
	Email    string `param:"email" validate:"required,email" example:"hi@example.com"`
	Password string `json:"password" validate:"required,cPassword"`
}

//func (uh *UserHandler) Login(ctx echo.Context) error {
//	req, err := utils.ValidateStruct(ctx, uh.validator, new(userLoginRequest))
//	if err != nil {
//		slog.Info("Http error", err.Error())
//		return err
//	}
//}
