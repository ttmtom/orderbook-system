package port

import (
	"github.com/labstack/echo/v4"
	"orderbook/internal/core/model"
)

type UserRepository interface {
	IsUserExist(email string) bool
	CreateUser(user *model.User) (*model.User, error)
	GetUserByIdHash(id string) (*model.User, error)
	GetUserById(id uint) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	UpdateUserLoginAt(user *model.User)
	UpdateUserLastAccessAt(userId string)
}

type UserService interface {
	UserRegistration(email string, password string) (*model.User, error)
	GetUserInformation(id string) (*model.User, error)
	GetUserById(id uint) (*model.User, error)
}

type UserController interface {
	Register(ctx echo.Context) error
	GetMe(ctx echo.Context) error
}
