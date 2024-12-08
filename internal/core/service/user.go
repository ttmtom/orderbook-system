package service

import (
	"errors"
	"log/slog"
	"orderbook/internal/core/model"
	"orderbook/internal/core/port"
	"orderbook/internal/pkg/security"
)

type UserEventTopic string

const (
	UserRegistrationSuccess UserEventTopic = "user-registration-success"
)

type UserServiceError string

const (
	EmailAlreadyExist UserServiceError = "EmailAlreadyExist"
	UserNotFound      UserServiceError = "UserNotFound"
	Unauthorized      UserServiceError = "Unauthorized"
	Unexpected        UserServiceError = "Unexpected"
)

type UserService struct {
	repo         port.UserRepository
	eventManager port.EventRepository
}

func NewUserService(
	resp port.UserRepository,
	eventManager port.EventRepository,
) port.UserService {
	return &UserService{
		resp,
		eventManager,
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

	err = us.eventManager.PublishEvent(string(UserRegistrationSuccess), event)
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

func (us *UserService) GetUserById(id uint) (*model.User, error) {
	user, err := us.repo.GetUserById(id)
	if err != nil {
		return nil, errors.New(string(UserNotFound))
	}
	return user, nil
}
