package service

import (
	"errors"
	"orderbook/internal/adapter/database/postgres/repository"
	"orderbook/internal/core/model"
)

type UserServiceError string

const (
	EmailAlreadyExist UserServiceError = "EmailAlreadyExist"
	UserNotFound      UserServiceError = "UserNotFound"
)

type UserService struct {
	resp *repository.UserRepository
}

func NewUserService(resp *repository.UserRepository) *UserService {
	return &UserService{
		resp,
	}
}

func (us *UserService) UserRegistration(email string) (*model.User, error) {
	if us.resp.IsUserExist(email) {
		return nil, errors.New(string(EmailAlreadyExist))
	}

	user := &model.User{
		Email: email,
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
