package users

import (
	"log"
	"orderbook-system/src/modules/database"
)

type Service struct {
	Database *database.Database
}

func (s *Service) CreateUser(user *User) error {
	log.Printf("Creating user: %v", user)
	return nil
}

func NewUserService(db *database.Database) *Service {
	return &Service{Database: db}
}
