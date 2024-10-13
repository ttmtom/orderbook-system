package orders

import (
	"errors"
	"github.com/google/uuid"
	"github.com/tarantool/go-tarantool"
	"log"
	"orderbook-system/src/modules/database"
	"orderbook-system/src/modules/users"
)

type Service struct {
	Database     *database.Database
	UsersService *users.Service
}

func (s *Service) canPlaceOrder(equity, margin, price, size float64) bool {
	newNotional := size * price
	newMargin := equity / (newNotional + (equity / margin))
	return newMargin > 0.1
}

func (s *Service) PlaceOrder(createOrderReq CreateOrderRequest) (*Order, error) {
	_, err := s.UsersService.GetUserById(createOrderReq.UserId)
	if err != nil {
		return &Order{}, err
	}

	price, err := s.UsersService.GetCurrentMarketPrice(createOrderReq.Market)
	if err != nil {
		return nil, err
	}

	equity, eErr := s.UsersService.CalculateAccountEquity(createOrderReq.UserId)
	if eErr != nil {
		return nil, eErr
	}
	margin, mErr := s.UsersService.CalculateAccountMargin(createOrderReq.UserId)
	if mErr != nil {
		return nil, mErr
	}
	if !s.canPlaceOrder(equity,
		margin,
		price,
		createOrderReq.Size,
	) {
		return nil, errors.New("insufficient margin")
	}

	order := Order{
		ID:     uuid.New().String(),
		UserID: createOrderReq.UserId,
		Market: createOrderReq.Market,
		Side:   createOrderReq.Side,
		Price:  price,
		Size:   createOrderReq.Size,
		Status: "pending",
	}

	s.handleOrderAsync(order)

	return &order, nil
}

func (s *Service) insertOrder(order Order) {

	var futures []*tarantool.Future

	tuples := [][]interface{}{
		{order.ID, order.UserID, order.Market, order.Side, order.Price, order.Size, order.Status},
	}
	for _, tuple := range tuples {
		request := tarantool.NewInsertRequest("orders").Tuple(tuple)
		futures = append(futures, s.Database.Connection.Do(request))
	}
}

func (s *Service) updateOrder(order Order) {

	var futures []*tarantool.Future

	tuples := [][]interface{}{
		{order.ID, order.UserID, order.Market, order.Side, order.Price, order.Size, order.Status},
	}
	for _, tuple := range tuples {
		request := tarantool.NewReplaceRequest("orders").Tuple(tuple)
		futures = append(futures, s.Database.Connection.Do(request))
	}
}

func (s *Service) handleOrderAsync(order Order) {
	go func() {
		s.insertOrder(order)
		s.matchOrders()
	}()
}

func (s *Service) getOrdersById(id string) ([]Order, error) {
	var orders []Order
	conn := s.Database.Connection
	err := conn.Do(
		tarantool.
			NewSelectRequest("orders").
			Limit(1).
			Iterator(tarantool.IterEq).
			Key([]interface{}{id}),
	).GetTyped(&orders)
	if err != nil {
		log.Fatalf("Failed to get orders by Side: %s", err)
		return nil, err
	}
	return orders, nil
}

func (s *Service) getOrdersBySide(side, status string) []Order {
	var orders []Order
	conn := s.Database.Connection
	err := conn.Do(
		tarantool.
			NewSelectRequest("orders").
			Index("todo_orders").
			Iterator(tarantool.IterEq).
			Key([]interface{}{side, status}),
	).GetTyped(&orders)
	if err != nil {
		log.Fatalf("Failed to get orders by Side: %s", err)
	}
	return orders
}

func (s *Service) matchOrders() {
	buyOrders := s.getOrdersBySide("buy", "pending")
	sellOrders := s.getOrdersBySide("sell", "pending")

	for _, buyOrder := range buyOrders {
		for _, sellOrder := range sellOrders {
			if buyOrder.Market == sellOrder.Market && buyOrder.Price >= sellOrder.Price {
				matchedSize := min(buyOrder.Size, sellOrder.Size)
				buyOrder.Size -= matchedSize
				sellOrder.Size -= matchedSize

				if buyOrder.Size == 0 {
					buyOrder.Status = "filled"
				}
				if sellOrder.Size == 0 {
					sellOrder.Status = "filled"
				}

				s.updateOrder(buyOrder)
				s.updateOrder(sellOrder)

				buyerPosition := users.Position{
					UserID:     buyOrder.UserID,
					Market:     buyOrder.Market,
					Side:       "long",
					EntryPrice: buyOrder.Price,
					Size:       matchedSize,
				}

				s.UsersService.UpdatePosition(buyerPosition)

				sellerPosition := users.Position{
					UserID:     sellOrder.UserID,
					Market:     sellOrder.Market,
					Side:       "short",
					EntryPrice: sellOrder.Price,
					Size:       matchedSize,
				}
				s.UsersService.UpdatePosition(sellerPosition)
			}
		}
	}
}

func NewOrdersService(db *database.Database, usersService *users.Service) *Service {
	return &Service{
		Database:     db,
		UsersService: usersService,
	}
}
