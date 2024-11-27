package middleware

import (
	"errors"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"orderbook/internal/core/service"
	"orderbook/internal/pkg/response"
	"orderbook/internal/pkg/security"
)

type AuthMiddleware struct {
	commonService *service.CommonService
	userService   *service.UserService
}

func NewAuthMiddleware(commonService *service.CommonService, userService *service.UserService) *AuthMiddleware {
	return &AuthMiddleware{commonService, userService}
}

func (am *AuthMiddleware) Handler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		accessToken, err := ctx.Cookie("x-access-token")
		var token *security.UserClaims
		if accessToken != nil {
			token, err = security.ValidateJwtToken(accessToken.Value)
		}
		var refreshToken *http.Cookie
		if err != nil {
			refreshToken, err = ctx.Cookie("x-refresh-token")
			if refreshToken == nil {
				return response.FailureResponse(http.StatusUnauthorized, errors.New("missing Token").Error())
			}
		}

		if refreshToken != nil {
			refreshToken, err := security.ValidateJwtToken(refreshToken.Value)
			if err != nil {
				return response.FailureResponse(http.StatusUnauthorized, err.Error())
			}
			refreshTokenTimeLimit, err := am.commonService.GetAccessTokenTimeLimit()
			if err != nil {
				return response.FailureResponse(http.StatusUnauthorized, err.Error())
			}

			user, err := am.userService.GetUserInformation(refreshToken.UserID)
			if err != nil {
				slog.Info("Failed to get user", err)
				return response.FailureResponse(http.StatusUnauthorized, err.Error())
			}

			newAccessToken, newToken, err := security.GenerateJwtToken(user, refreshTokenTimeLimit)
			if err != nil {
				slog.Info("Failed to gen access time limit", err)
				return response.FailureResponse(http.StatusUnauthorized, err.Error())
			}
			token = newToken
			response.SetSecureCookies(
				ctx,
				"x-access-token",
				*newAccessToken,
				newToken.ExpiresAt.Time,
				int(newToken.MaxAge),
			)
		}

		ctx.Set("user-claims", token)

		return next(ctx)
	}
}
