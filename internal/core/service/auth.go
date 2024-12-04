package service

import (
	"errors"
	"log/slog"
	"orderbook/internal/adapter/database/postgres/repository"
	"orderbook/internal/core/model"
	"orderbook/internal/pkg/security"
)

type AuthService struct {
	commonService  *CommonService
	userService    *UserService
	userRepository *repository.UserRepository
}

func NewAuthService(
	commonService *CommonService,
	userService *UserService,
	userRepository *repository.UserRepository,
) *AuthService {
	return &AuthService{
		commonService,
		userService,
		userRepository,
	}
}

type UserLoginToken struct {
	AccessToken        string               `json:"accessToken"`
	AccessTokenClaims  *security.UserClaims `json:"accessTokenClaims"`
	RefreshToken       string               `json:"refreshToken"`
	RefreshTokenClaims *security.UserClaims `json:"refreshTokenClaims"`
}

func (as *AuthService) generateUserLoginToken(user *model.User) (*UserLoginToken, error) {
	timeLimits, err := as.commonService.GetJwtTokenTimeLimit()
	if err != nil {
		slog.Info("Failed to get jwt token time limit", err)
		return nil, errors.New(string(Unexpected))
	}

	accessToken, accessClaims, err := security.GenerateJwtToken(user, timeLimits.AccessTokenDuration, security.AccessToken)
	if err != nil {
		slog.Info("Failed to gen access time limit", err)
		return nil, errors.New(string(Unexpected))
	}
	refreshToken, refreshClaims, err := security.GenerateJwtToken(user, timeLimits.RefreshTokenDuration, security.AccessToken)
	if err != nil {
		slog.Info("Failed to gen refresh time limit", err)
		return nil, errors.New(string(Unexpected))
	}

	return &UserLoginToken{
		*accessToken,
		accessClaims,
		*refreshToken,
		refreshClaims,
	}, nil
}

func (as *AuthService) UserLogin(email string, password string) (*model.User, *UserLoginToken, error) {
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
		return nil, nil, errors.New(string(Unauthorized))
	}

	as.userRepository.UpdateUserLoginAt(user)

	return user, jwt, nil
}

func (as *AuthService) UserAccess(user *security.UserClaims) {
	as.userRepository.UpdateUserLastAccessAt(user.UserID)
}

func (as *AuthService) RefreshToken(token string) (*UserLoginToken, error) {
	userClaims, err := security.ValidateJwtToken(token, security.RefreshToken)
	if err != nil {
		return nil, err
	}

	user, err := as.userRepository.GetUserByIdHash(userClaims.UserID)

	timeLimits, err := as.commonService.GetJwtTokenTimeLimit()
	if err != nil {
		slog.Info("Failed to get jwt token time limit", err)
		return nil, errors.New(string(Unexpected))
	}

	accessToken, accessClaims, err := security.GenerateJwtToken(user, timeLimits.AccessTokenDuration, security.AccessToken)
	if err != nil {
		slog.Info("Failed to gen access time limit", err)
		return nil, errors.New(string(Unexpected))
	}

	return &UserLoginToken{
		AccessToken:        *accessToken,
		AccessTokenClaims:  accessClaims,
		RefreshToken:       token,
		RefreshTokenClaims: userClaims,
	}, nil
}
