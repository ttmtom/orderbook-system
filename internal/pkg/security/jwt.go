package security

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"orderbook/internal/core/model"
	"os"
	"time"
)

type Claims struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	DisplayName string `json:"displayName"`
	UpdatedAt   string `json:"updatedAt"`
	jwt.RegisteredClaims
}

func GenerateUserToken(user *model.User, expiration time.Duration) (string, error) {
	claims := &Claims{
		ID: user.IDHash,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(os.Getenv("APP_SECRET_KEY")))
	if err != nil {
		println(err.Error())
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return signedToken, nil
}
