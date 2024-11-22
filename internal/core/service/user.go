package service

import (
	"errors"
	"orderbook/internal/adapter/database/postgres/repository"
	"orderbook/internal/core/model"
	"orderbook/internal/pkg/security"
)

type UserServiceError string

const (
	EmailAlreadyExist UserServiceError = "EmailAlreadyExist"
	UserNotFound      UserServiceError = "UserNotFound"
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

func (us *UserService) GetUserInformation(id uint) (*model.User, error) {
	user, err := us.resp.GetUserById(id)
	if err != nil {
		return nil, errors.New(string(UserNotFound))
	}

	return user, nil
}
