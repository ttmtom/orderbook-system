package users

import (
	"net/http"
	"orderbook-system/src/modules/database"
)

type Module struct {
	UsersController *Controller
	UserService     *Service
}

func (m *Module) Router() http.Handler {
	return m.UsersController.Router
}

func NewUsersModule(db *database.Database) (*Module, error) {
	service := NewUserService(db)
	return &Module{
		UsersController: NewUserController(service),
		UserService:     service,
	}, nil
}
