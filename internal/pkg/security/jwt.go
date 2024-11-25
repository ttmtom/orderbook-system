package security

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"orderbook/internal/core/model"
	"os"
	"time"
)

type Claims struct {
	ID     string `json:"id"`
	Email  string `json:"email"`
	MaxAge uint   `json:"maxAge"`
	jwt.RegisteredClaims
}

func GenerateJwtToken(user *model.User, expiration uint) (*string, *Claims, error) {
	claims := &Claims{
		ID:     user.IDHash,
		Email:  user.Email,
		MaxAge: expiration / 1000,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expiration))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(os.Getenv("APP_SECRET_KEY")))
	if err != nil {
		println(err.Error())
		return nil, nil, fmt.Errorf("failed to sign token: %w", err)
	}

	return &signedToken, claims, nil
}
