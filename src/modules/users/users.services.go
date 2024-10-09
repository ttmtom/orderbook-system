package users

import (
	"github.com/tarantool/go-tarantool"
	"log"
	"orderbook-system/src/modules/database"
)

type Service struct {
	Database *database.Database
}

func (s *Service) CreateUser(user *User) error {
	log.Printf("Creating user: %v", user)
	data, err := s.Database.Connection.Do(
		tarantool.NewSelectRequest("users")).Get()
	if err != nil {
		log.Println("Error:", err)
	} else {
		log.Println("Data:", data)
	}
	return nil
}

func (s *Service) CreateAccount(account Account) error {
	log.Printf("Creating account: %v", account)
	data, err := s.Database.Connection.Do(
		tarantool.NewSelectRequest("accounts")).Get()
	if err != nil {
		log.Println("Error:", err)
	} else {
		log.Println("Data:", data)
	}
	return nil
}

func NewUserService(db *database.Database) *Service {
	return &Service{Database: db}
}
