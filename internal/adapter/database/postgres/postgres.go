package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"orderbook/config"
)

type Postgres struct {
	DB *gorm.DB
}

func New(config config.DatabaseConfig) (*Postgres, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Host,
		config.User,
		config.Password,
		config.DBName,
		config.Port,
	)

	PostgresDb, dbOpenErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if dbOpenErr != nil {
		return nil, dbOpenErr
	}
	return &Postgres{
		DB: PostgresDb,
	}, nil
}
