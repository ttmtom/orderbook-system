package service

import (
	"errors"
	"log/slog"
	"orderbook/internal/adapter/database/postgres/repository"
	"orderbook/internal/core/model"
	"orderbook/internal/pkg/security"
)

type UserServiceError string

const (
	EmailAlreadyExist UserServiceError = "EmailAlreadyExist"
	UserNotFound      UserServiceError = "UserNotFound"
	Unauthorized      UserServiceError = "Unauthorized"
	Unexpected        UserServiceError = "Unexpected"
)

type UserService struct {
	resp          *repository.UserRepository
	commonService *CommonService
}

func NewUserService(resp *repository.UserRepository, commonService *CommonService) *UserService {
	return &UserService{
		resp,
		commonService,
	}
}

func (us *UserService) UserRegistration(email string, password string) (*model.User, error) {
	if us.resp.IsUserExist(email) {
		return nil, errors.New(string(EmailAlreadyExist))
	}

	hashedPWD, err := security.HashPassword(password)
	if err != nil {
		return nil, errors.New(string(Unexpected))
	}

	user := &model.User{
		Email:        email,
		IDHash:       security.HashEmail(email),
		PasswordHash: hashedPWD,
	}

	return us.resp.CreateUser(user)
}

func (us *UserService) GetUserInformation(id string) (*model.User, error) {
	user, err := us.resp.GetUserByIdHash(id)
	if err != nil {
		return nil, errors.New(string(UserNotFound))
	}

	return user, nil
}

type UserLoginToken struct {
	AccessToken        string               `json:"accessToken"`
	AccessTokenClaims  *security.UserClaims `json:"accessTokenClaims"`
	RefreshToken       string               `json:"refreshToken"`
	RefreshTokenClaims *security.UserClaims `json:"refreshTokenClaims"`
}

func (us *UserService) generateUserLoginToken(user *model.User) (*UserLoginToken, error) {
	timeLimits, err := us.commonService.GetJwtTokenTimeLimit()
	if err != nil {
		slog.Info("Failed to get jwt token time limit", err)
		return nil, errors.New(string(Unexpected))
	}

	accessToken, accessClaims, err := security.GenerateJwtToken(user, timeLimits.AccessTokenDuration)
	if err != nil {
		slog.Info("Failed to gen access time limit", err)
		return nil, errors.New(string(Unexpected))
	}
	refreshToken, refreshClaims, err := security.GenerateJwtToken(user, timeLimits.RefreshTokenDuration)
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

func (us *UserService) UserLogin(email string, password string) (*model.User, *UserLoginToken, error) {
	user, err := us.resp.GetUserByEmail(email)
	if err != nil {
		slog.Info("Email not found %s", email)
		return nil, nil, errors.New(string(Unauthorized))
	}

	err = security.ComparePassword(password, user.PasswordHash)
	if err != nil {
		slog.Info("Password not match %s", email)
		return nil, nil, errors.New(string(Unauthorized))
	}

	jwt, err := us.generateUserLoginToken(user)
	if err != nil {
		slog.Info("Failed to generate token %s", email)
		return nil, nil, errors.New(string(Unauthorized))
	}

	return user, jwt, nil
}
