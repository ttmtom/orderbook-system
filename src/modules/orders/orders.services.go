package orders

import (
	"orderbook-system/src/modules/database"
	"orderbook-system/src/modules/users"
)

type Service struct {
	Database     *database.Database
	UsersService *users.Service
}

func NewOrdersService(db *database.Database, usersService *users.Service) *Service {
	return &Service{
		Database:     db,
		UsersService: usersService,
	}
}
