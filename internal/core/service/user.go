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
