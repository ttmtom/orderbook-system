//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"orderbook-system/src/utils/database"
)

func InitializeDBConnection() (*database.DBConnection, error) {
	wire.Build(database.NewConnection)
	return &database.DBConnection{}, nil
}
