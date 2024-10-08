//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"orderbook-system/src/modules/database"
	"orderbook-system/src/modules/users"
)

func InitializeDatabase() (*database.Database, error) {
	wire.Build(database.NewConnection)
	return &database.Database{}, nil
}

func InitializeUserModule(db *database.Database) (*users.Module, error) {
	wire.Build(users.NewUsersModule)
	return &users.Module{}, nil
}
