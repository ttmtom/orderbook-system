package users

import (
	"github.com/gorilla/mux"
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

func NewUsersModule(router *mux.Router, db *database.Database) (*Module, error) {
	service := NewUserService(db)
	return &Module{
		UsersController: NewUserController(router, service),
		UserService:     service,
	}, nil
}
