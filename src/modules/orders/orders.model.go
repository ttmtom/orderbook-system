package orders

type Rate struct {
	Rate      float64
	Currency  string
	UpdatedAt int64
}

type Order struct {
	ID        int
	UserID    string
	Rate      Rate
	Amount    float64
	Type      string
	Status    string
	CreatedAt int64
	UpdatedAt int64
	Pair      int
}
