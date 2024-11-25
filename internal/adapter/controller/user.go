package controller

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
	ID          string    `json:"id"`
	Email       string    `json:"email"`
	DisplayName *string   `json:"displayName"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (uc *UserController) formatUserResponse(user *model.User) *useResponse {
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

type getUserRequest struct {
	ID string `param:"idHash" validate:"required,min=1" example:"1"`
}

func (uc *UserController) GetUser(ctx echo.Context) error {
	req, err := utils.ValidateStruct(ctx, uc.validator, new(getUserRequest))
	if err != nil {
		slog.Info("Validation Error", err.Error())
		return err
	}

	user, err := uc.svc.GetUserInformation(req.ID)
	if err != nil {
		slog.Info("Error during registration", err)

		if err.Error() == string(service.UserNotFound) {
			return response.FailureResponse(http.StatusNotFound, err.Error())
		}
		return response.FailureResponse(http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(ctx, http.StatusOK, uc.formatUserResponse(user))
}

type userLoginRequest struct {
	Email    string `param:"email" validate:"required,email" example:"hi@example.com"`
	Password string `json:"password" validate:"required,cPassword"`
}

func (uc *UserController) Login(ctx echo.Context) error {
	req, err := utils.ValidateStruct(ctx, uc.validator, new(userLoginRequest))
	if err != nil {
		slog.Info("Validation Error", err.Error())
		return err
	}

	user, jwt, err := uc.svc.UserLogin(req.Email, req.Password)
	if err != nil {
		slog.Info("Error during login", err)

		if err.Error() == string(service.Unauthorized) {
			return response.FailureResponse(http.StatusNotFound, err.Error())
		}
		return response.FailureResponse(http.StatusInternalServerError, err.Error())
	}

	ctx.SetCookie(&http.Cookie{
		Name:     "x-access-token",
		Value:    jwt.AccessToken,
		Expires:  jwt.AccessTokenClaims.ExpiresAt.Time,
		MaxAge:   int(jwt.AccessTokenClaims.MaxAge),
		Secure:   true,
		HttpOnly: true,
		SameSite: 3,
	})
	ctx.SetCookie(&http.Cookie{
		Name:     "x-refresh-token",
		Value:    jwt.RefreshToken,
		Expires:  jwt.RefreshTokenClaims.ExpiresAt.Time,
		MaxAge:   int(jwt.RefreshTokenClaims.MaxAge),
		Secure:   true,
		HttpOnly: true,
		SameSite: 3,
	})

	return response.SuccessResponse(ctx, http.StatusOK, uc.formatUserResponse(user))
}
