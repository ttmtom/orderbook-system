package service

import (
	"errors"
	"log/slog"
	"orderbook/internal/core/model"
	"orderbook/internal/core/port"
	"orderbook/internal/pkg/security"
)

type AuthService struct {
	commonRepository port.CommonRepository
	userRepository   port.UserRepository
}

func NewAuthService(
	commonRepository port.CommonRepository,
	userRepository port.UserRepository,
) port.AuthService {
	return &AuthService{
		commonRepository,
		userRepository,
	}
}

func (as *AuthService) getTimeLimit(id string, c chan *model.TimeLimit) {
	limit, err := as.commonRepository.GetTimeLimit(id)
	if err != nil {
		slog.Info("error on getting time limit", "id", id, "err", err)
	} else {
		c <- limit
	}
	close(c)
}

func (as *AuthService) generateUserLoginToken(user *model.User) (*port.UserLoginToken, error) {
	accessChannel := make(chan *model.TimeLimit)
	refreshChannel := make(chan *model.TimeLimit)

	go as.getTimeLimit("access_token", accessChannel)
	go as.getTimeLimit("refresh_token", refreshChannel)

	accessTimeLimit := <-accessChannel
	refreshTimeLimit := <-refreshChannel
	if accessTimeLimit == nil || refreshTimeLimit == nil {
		slog.Info("Failed to get access time limit")
		return nil, errors.New(string(Unexpected))
	}

	accessToken, accessClaims, err := security.GenerateJwtToken(user.IDHash, "user", accessTimeLimit.Time, security.AccessToken)
	if err != nil {
		slog.Info("Failed to gen access time limit", err)
		return nil, errors.New(string(Unexpected))
	}
	refreshToken, refreshClaims, err := security.GenerateJwtToken(user.IDHash, "user", refreshTimeLimit.Time, security.RefreshToken)
	if err != nil {
		slog.Info("Failed to gen refresh time limit", err)
		return nil, errors.New(string(Unexpected))
	}

	return &port.UserLoginToken{
		AccessToken:        *accessToken,
		AccessTokenClaims:  accessClaims,
		RefreshToken:       *refreshToken,
		RefreshTokenClaims: refreshClaims,
	}, nil
}

func (as *AuthService) UserLogin(email string, password string) (*model.User, *port.UserLoginToken, error) {
	user, err := as.userRepository.GetUserByEmail(email)
	if err != nil {
		slog.Info("Email not found %s", email)
		return nil, nil, errors.New(string(Unauthorized))
	}

	err = security.ComparePassword(password, user.PasswordHash)
	if err != nil {
		slog.Info("Password not match %s", email)
		return nil, nil, errors.New(string(Unauthorized))
	}

	jwt, err := as.generateUserLoginToken(user)
	if err != nil {
		slog.Info("Failed to generate token %s", email)
		return nil, nil, err
	}

	as.userRepository.UpdateUserLoginAt(user)

	return user, jwt, nil
}

func (as *AuthService) UserAccess(user *security.UserClaims) {
	as.userRepository.UpdateUserLastAccessAt(user.UserID)
}

func (as *AuthService) RefreshToken(accessToken string, refreshToken string) (*port.UserLoginToken, error) {
	userRefreshClaims, err := security.ValidateJwtToken(refreshToken, security.RefreshToken)
	if err != nil {
		return nil, err
	}

	accessClaims, err := security.ValidateJwtToken(accessToken, security.AccessToken, true)
	if err != nil {
		return nil, err
	}

	if accessClaims.UserID != userRefreshClaims.UserID {
		return nil, err
	}

	user, err := as.userRepository.GetUserByIdHash(userRefreshClaims.UserID)

	limit, err := as.commonRepository.GetTimeLimit("access_token")
	if err != nil {
		slog.Info("Failed to get jwt token time limit", err)
		return nil, errors.New(string(Unexpected))
	}

	newAccessToken, newAccessClaims, err := security.GenerateJwtToken(user.IDHash, "user", limit.Time, security.AccessToken)
	if err != nil {
		slog.Info("Failed to gen access time limit", err)
		return nil, errors.New(string(Unexpected))
	}

	return &port.UserLoginToken{
		AccessToken:        *newAccessToken,
		AccessTokenClaims:  newAccessClaims,
		RefreshToken:       refreshToken,
		RefreshTokenClaims: userRefreshClaims,
	}, nil
}
