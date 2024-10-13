package orders

type Order struct {
	ID     string
	UserID string
	Market string
	Side   string
	Price  float64
	Size   float64
	Status string
}

type Position struct {
	UserID     string
	Market     string
	Side       string
	EntryPrice float64
	Size       float64
}
