package orders

import (
	"github.com/gorilla/mux"
	"orderbook-system/src/modules/database"
	"orderbook-system/src/modules/users"
)

type Module struct {
	OrdersController *Controller
	OrdersService    *Service
	UsersService     *users.Service
}

func NewOrdersModule(router *mux.Router, db *database.Database, usersService *users.Service) (*Module, error) {
	server := NewOrdersService(db, usersService)

	return &Module{
		OrdersController: NewOrdersController(router, server),
		OrdersService:    server,
	}, nil
}
