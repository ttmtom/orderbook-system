package service

import (
	"errors"
	"log/slog"
	"orderbook/internal/adapter/database/postgres/repository"
	"orderbook/internal/core/model"
	"orderbook/internal/pkg/security"
	"time"
)

type UserServiceError string

const (
	EmailAlreadyExist UserServiceError = "EmailAlreadyExist"
	UserNotFound      UserServiceError = "UserNotFound"
	Unauthorized      UserServiceError = "Unauthorized"
	Unexpected        UserServiceError = "Unexpected"
)

type UserService struct {
	resp *repository.UserRepository
}

func NewUserService(resp *repository.UserRepository) *UserService {
	return &UserService{
		resp,
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

func (us *UserService) UserLogin(email string, password string) (*model.User, *string, error) {
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

	jwt, err := security.GenerateUserToken(user, time.Duration(15*60*1000))
	if err != nil {
		slog.Info("Failed to generate token %s", email)
		return nil, nil, errors.New(string(Unauthorized))
	}

	return user, &jwt, nil

}
