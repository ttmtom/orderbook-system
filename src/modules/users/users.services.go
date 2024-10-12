package users

import (
	"errors"
	"fmt"
	"github.com/tarantool/go-tarantool"
	"log"
	"orderbook-system/src/modules/database"
)

type Service struct {
	Database *database.Database
}

func (s *Service) getUserById(id string) (*User, error) {
	log.Printf("Getting user by ID: %v", id)
	conn := s.Database.Connection
	var users []User
	err := conn.Do(
		tarantool.
			NewSelectRequest("users").
			Limit(1).
			Iterator(tarantool.IterEq).
			Key([]interface{}{id}),
	).GetTyped(&users)

	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, nil
	}

	log.Printf("User found: %v", users[0])

	return &users[0], nil
}

func (s *Service) getAccountById(userId string) (*Account, error) {
	log.Printf("Getting account by ID: %v", userId)
	conn := s.Database.Connection
	var accounts []Account
	err := conn.Do(
		tarantool.
			NewSelectRequest("accounts").
			Limit(5).
			Iterator(tarantool.IterEq).
			Key([]interface{}{userId}),
	).GetTyped(&accounts)
	if err != nil {
		return nil, err
	}
	if len(accounts) == 0 {
		return nil, nil
	}

	return &accounts[0], nil
}

func (s *Service) CreateUser(user *User) error {
	log.Printf("Creating user: %v", user)

	exist, _ := s.getUserById(user.ID)

	if exist != nil {
		return errors.New("user already exists")
	}

	var futures []*tarantool.Future

	tuples := [][]interface{}{
		{user.ID},
	}
	for _, tuple := range tuples {
		request := tarantool.NewInsertRequest("users").Tuple(tuple)
		futures = append(futures, s.Database.Connection.Do(request))
	}

	log.Printf("Created user: %v", futures)
	return nil
}

func (s *Service) CreateAccounts(accounts []Account) error {
	log.Printf("Creating account: %v", accounts)
	conn := s.Database.Connection

	var futures []*tarantool.Future
	for _, acc := range accounts {
		log.Printf("Creating account: %v", acc)
		request := tarantool.NewInsertRequest("accounts").Tuple([]interface{}{
			acc.UserID, acc.Amount, acc.Currency,
		})
		futures = append(futures, conn.Do(request))
	}
	fmt.Println("Inserted tuples:")
	for _, future := range futures {
		result, err := future.Get()
		if err != nil {
			fmt.Println("Got an error:", err)
		} else {
			fmt.Println(result)
		}
	}
	return nil
}

func (s *Service) GetUserAccounts(userId string) ([]Account, error) {
	log.Printf("Getting accounts for user with ID: %v", userId)
	conn := s.Database.Connection
	var accounts []Account
	err := conn.Do(
		tarantool.
			NewSelectRequest("accounts").
			Index("user_id").
			Limit(5).
			Iterator(tarantool.IterEq).
			Key([]interface{}{userId}),
	).GetTyped(&accounts)

	log.Printf("Accounts found: %v", accounts)
	log.Printf("Error: %v", err)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func NewUserService(db *database.Database) *Service {
	return &Service{Database: db}
}
