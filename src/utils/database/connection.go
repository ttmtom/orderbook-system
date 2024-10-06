package database

import (
	"github.com/tarantool/go-tarantool"
	"log"
	"os"
)

type DBConnection struct {
	Connection *tarantool.Connection
}

func NewConnection() (*DBConnection, error) {
	opts := tarantool.Opts{
		User: os.Getenv("DATABASE_USER"),
		Pass: os.Getenv("DATABASE_PASSWORD"),
	}

	log.Printf("Connecting to tarantool at %v\n", os.Getenv("TARANTOOL_URL"))

	conn, err := tarantool.Connect(os.Getenv("TARANTOOL_URL"), opts)
	if err != nil {
		log.Fatalf("Failed to connect to tarantool: %v\n", err)
		return nil, err
	}
	return &DBConnection{
		Connection: conn,
	}, nil
}
