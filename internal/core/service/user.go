package service

import (
	"errors"
	"log/slog"
	"orderbook/internal/adapter/database/postgres/repository"
	"orderbook/internal/adapter/kafka"
	"orderbook/internal/core/model"
	"orderbook/internal/pkg/security"
	"strings"
)

type UserServiceError string

const (
	EmailAlreadyExist UserServiceError = "EmailAlreadyExist"
	UserNotFound      UserServiceError = "UserNotFound"
	Unauthorized      UserServiceError = "Unauthorized"
	Unexpected        UserServiceError = "Unexpected"
)

type UserService struct {
	repo          *repository.UserRepository
	commonService *CommonService
	kafkaManager  *kafka.Manager
}

func NewUserService(
	resp *repository.UserRepository,
	commonService *CommonService,
	kafkaManager *kafka.Manager,
) *UserService {

	return &UserService{
		resp,
		commonService,
		kafkaManager,
	}
}

type UserRegistrationSuccessEvent struct {
	ID uint `json:"id"`
}

func (us *UserService) UserRegistration(email string, password string) (*model.User, error) {
	if us.repo.IsUserExist(email) {
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

	user, err = us.repo.CreateUser(user)
	if err != nil {
		return nil, errors.New(string(Unexpected))
	}

	event := UserRegistrationSuccessEvent{
		ID: user.ID,
	}

	topic := strings.Join([]string{"user", "registration", "success"}, "-")

	err = us.kafkaManager.PublishEvent(topic, event)
	if err != nil {
		slog.Error("Failed to marshal event data", "error", err)
		return nil, errors.New(string(Unexpected))
	}
	return user, nil
}

func (us *UserService) GetUserInformation(id string) (*model.User, error) {
	user, err := us.repo.GetUserByIdHash(id)
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

func (us *UserService) UserLogin(email string, password string) (*model.User, *UserLoginToken, error) {
	user, err := us.repo.GetUserByEmail(email)
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

	us.repo.UpdateUserLoginAt(user)

	return user, jwt, nil
}

func (us *UserService) UserAccess(user *security.UserClaims) {
	us.repo.UpdateUserLastAccessAt(user.UserID)
}

func (us *UserService) RefreshToken(token string) (*UserLoginToken, error) {
	userClaims, err := security.ValidateJwtToken(token, security.RefreshToken)
	if err != nil {
		return nil, err
	}

	user, err := us.repo.GetUserByIdHash(userClaims.UserID)

	timeLimits, err := us.commonService.GetJwtTokenTimeLimit()
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
