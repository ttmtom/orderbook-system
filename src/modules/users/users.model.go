package users

type User struct {
	ID      string
	Balance float64
}

type Position struct {
	UserID     string
	Market     string
	Side       string
	EntryPrice float64
	Size       float64
}
