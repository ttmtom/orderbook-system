package controller

import (
	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"orderbook/internal/core/model"
	"orderbook/internal/core/service"
	"orderbook/internal/pkg/response"
	"orderbook/internal/pkg/security"
	"orderbook/pkg/utils"
	"time"
)

type UserController struct {
	svc       *service.UserService
	validator *validator.Validate
}

func NewUserController(validator *validator.Validate, svc *service.UserService) *UserController {
	return &UserController{
		svc,
		validator,
	}
}

type useResponse struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	DisplayName  *string   `json:"displayName"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	LastLoginAt  time.Time `json:"lastLoginAt"`
	LastAccessAt time.Time `json:"lastAccessAt"`
}

func (uc *UserController) formatUserResponse(user *model.User) *useResponse {
	return &useResponse{
		Email:        user.Email,
		ID:           user.IDHash,
		DisplayName:  user.DisplayName,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		LastLoginAt:  user.LastLoginAt,
		LastAccessAt: user.LastAccessAt,
	}
}

type registerRequest struct {
	Email    string `json:"email" validate:"required,email" example:"test@example.com"`
	Password string `json:"password" validate:"required,cPassword"`
}

func (uc *UserController) Register(ctx echo.Context) error {
	req, err := utils.ValidateStruct(ctx, uc.validator, new(registerRequest))
	if err != nil {
		slog.Info("Validation Error", err.Error())
		return err
	}

	user, err := uc.svc.UserRegistration(req.Email, req.Password)
	if err != nil {
		slog.Info("Error during registration", err)
		if err.Error() == string(service.EmailAlreadyExist) {
			return response.FailureResponse(http.StatusConflict, err.Error())
		}
		return response.FailureResponse(http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(ctx, http.StatusOK, uc.formatUserResponse(user))
}

func (uc *UserController) GetUser(ctx echo.Context) error {
	userToken := ctx.Get("user").(*jwt.Token)
	userClaims := userToken.Claims.(*security.UserClaims)

	user, err := uc.svc.GetUserInformation(userClaims.UserID)
	if err != nil {
		slog.Info("Error on get user info", err)

		if err.Error() == string(service.UserNotFound) {
			return response.FailureResponse(http.StatusNotFound, err.Error())
		}
		return response.FailureResponse(http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(ctx, http.StatusOK, uc.formatUserResponse(user))
}
