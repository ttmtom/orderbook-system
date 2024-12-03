package security

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log/slog"
	"orderbook/internal/core/model"
	"time"
)

type JWTTokenType string

const (
	AccessToken  JWTTokenType = "accessToken"
	RefreshToken JWTTokenType = "refreshToken"
)

type UserClaims struct {
	UserID string       `json:"id"`
	Email  string       `json:"email"`
	MaxAge uint         `json:"maxAge"`
	Type   JWTTokenType `json:"type"`
	jwt.RegisteredClaims
}

type JwtSecurity struct {
	secretKey []byte
}

var js *JwtSecurity

func InitJwtSecurity(secretKey string) {
	js = &JwtSecurity{
		secretKey: []byte(secretKey),
	}
}

func GenerateJwtToken(user *model.User, expiration uint, tokenType JWTTokenType) (*string, *UserClaims, error) {
	claims := &UserClaims{
		user.IDHash,
		user.Email,
		expiration,
		tokenType,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expiration) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(js.secretKey)
	if err != nil {
		println(err.Error())
		return nil, nil, fmt.Errorf("failed to sign token: %w", err)
	}

	return &signedToken, claims, nil
}

func ValidateJwtToken(tokenString string, tokenType JWTTokenType) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			slog.Info("Invalid token method")
			return nil, jwt.ErrSignatureInvalid
		}
		return js.secretKey, nil
	})

	user, ok := token.Claims.(*UserClaims)
	if !ok || !token.Valid || user.Type != tokenType {
		slog.Info("Error on Validate token", "err", ok, "token", token)
		return nil, err
	}
	return user, nil
}
