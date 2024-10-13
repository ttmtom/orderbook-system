package users

import (
	"errors"
	"github.com/tarantool/go-tarantool"
	"log"
	"orderbook-system/src/modules/database"
)

type Service struct {
	Database *database.Database
}

func (s *Service) GetUserById(id string) (*User, error) {
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

func (s *Service) CreateUser(user *User) error {
	log.Printf("Creating user: %v", user)

	exist, _ := s.GetUserById(user.ID)

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

func (s *Service) CalculateAccountEquity(userID string) float64 {
	user, _ := s.GetUserById(userID)
	positions := s.GetPositions(userID)

	unrealizedPNL := 0.0
	for _, position := range positions {
		currentPrice := s.GetCurrentMarketPrice(position.Market)
		unrealizedPNL += position.SideMultiplier() * position.Size * (currentPrice - position.EntryPrice)
	}

	accountEquity := user.Balance + unrealizedPNL
	return accountEquity
}

func (s *Service) GetPositions(userID string) []Position {
	var positions []Position
	conn := s.Database.Connection
	err := conn.SelectTyped("positions", "primary", 0, 100, tarantool.IterEq, []interface{}{userID}, &positions)
	if err != nil {
		log.Fatalf("Failed to get positions: %s", err)
	}
	return positions
}

func (s *Service) UpdatePosition(position Position) {
	var futures []*tarantool.Future

	tuples := [][]interface{}{
		{position.UserID, position.Market, position.Side, position.EntryPrice, position.Size},
	}
	for _, tuple := range tuples {
		request := tarantool.NewReplaceRequest("positions").Tuple(tuple)
		futures = append(futures, s.Database.Connection.Do(request))
	}
}

func (s *Service) GetCurrentMarketPrice(market string) float64 {
	return 0
}

func (s *Service) CalculateAccountMargin(userID string) float64 {
	accountEquity := s.CalculateAccountEquity(userID)
	positions := s.GetPositions(userID)

	totalPositionNotional := 0.0
	for _, position := range positions {
		currentPrice := s.GetCurrentMarketPrice(position.Market)
		totalPositionNotional += position.Size * currentPrice
	}

	accountMargin := accountEquity / totalPositionNotional
	return accountMargin
}

func (p *Position) SideMultiplier() float64 {
	if p.Side == "long" {
		return 1.0
	}
	return -1.0
}

func NewUserService(db *database.Database) *Service {
	return &Service{Database: db}
}
