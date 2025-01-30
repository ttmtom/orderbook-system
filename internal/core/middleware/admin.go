package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"orderbook/config"
	"orderbook/internal/core/port"
	"orderbook/internal/pkg/response"
	"orderbook/internal/pkg/security"
)

type AdminMiddleware struct {
	appConfig    config.AppConfig
	adminService port.AdminService
}

func NewAdminMiddleware() port.AdminMiddleware {
	return &AdminMiddleware{}
}

func (am *AdminMiddleware) AdminAuthHandler() func(next echo.HandlerFunc) echo.HandlerFunc {
	errorHandler := func(ctx echo.Context, err error) error {
		slog.Info("admin auth error", "err", err)
		return response.FailureResponse(http.StatusUnauthorized, map[string]string{
			"error":   "Invalid token",
			"message": err.Error(),
		})
	}

	mc := echojwt.Config{
		ErrorHandler: errorHandler,
		SigningKey:   []byte(am.appConfig.AdminSecurityKey),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(security.UserClaims)
		},
	}

	return echojwt.WithConfig(mc)
}
