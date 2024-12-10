package port

import (
	"github.com/labstack/echo/v4"
	"orderbook/internal/core/model"
	"orderbook/internal/pkg/security"
)

type UserLoginToken struct {
	AccessToken        string               `json:"accessToken"`
	AccessTokenClaims  *security.UserClaims `json:"accessTokenClaims"`
	RefreshToken       string               `json:"refreshToken"`
	RefreshTokenClaims *security.UserClaims `json:"refreshTokenClaims"`
}

type AuthController interface {
	Login(ctx echo.Context) error
	RefreshToken(ctx echo.Context) error
}

type AuthService interface {
	UserLogin(email string, password string) (*model.User, *UserLoginToken, error)
	UserAccess(user *security.UserClaims)
	RefreshToken(accessToken string, refreshToken string) (*UserLoginToken, error)
}

type AuthMiddleware interface {
	HeaderAuthHandler(args ...bool) func(next echo.HandlerFunc) echo.HandlerFunc
}
