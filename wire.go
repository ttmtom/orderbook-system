//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/gorilla/mux"
	"orderbook-system/src/modules/database"
	"orderbook-system/src/modules/orders"
	"orderbook-system/src/modules/users"
)

func InitializeDatabase() (*database.Database, error) {
	wire.Build(database.NewConnection)
	return &database.Database{}, nil
}

func InitializeUserModule(router *mux.Router, db *database.Database) (*users.Module, error) {
	wire.Build(users.NewUsersModule)
	return &users.Module{}, nil
}

func InitializeOrderModule(router *mux.Router, db *database.Database, usersService *users.Service) (*orders.Module, error) {
	wire.Build(orders.NewOrdersModule)
	return &orders.Module{}, nil
}
