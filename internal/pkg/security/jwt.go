package security

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log/slog"
	"time"
)

type JWTTokenType string

const (
	AccessToken  JWTTokenType = "accessToken"
	RefreshToken JWTTokenType = "refreshToken"
)

type UserClaims struct {
	UserID string       `json:"id"`
	Role   string       `json:"role"`
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

func GenerateJwtToken(signID string, role string, expiration uint, tokenType JWTTokenType) (*string, *UserClaims, error) {
	claims := &UserClaims{
		signID,
		role,
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

func ValidateJwtToken(tokenString string, tokenType JWTTokenType, args ...bool) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			slog.Info("Invalid token method")
			return nil, jwt.ErrSignatureInvalid
		}
		return js.secretKey, nil
	})

	ignoreError := false
	if len(args) > 0 {
		ignoreError = args[0]
	}

	user, ok := token.Claims.(*UserClaims)
	if !ignoreError && (!ok || !token.Valid || user.Type != tokenType) {
		slog.Info("Error on Validate token", "err", ok, "token", token)
		return nil, err
	}

	return user, nil
}
