package users

type User struct {
	ID string
}

type Account struct {
	UserID   string
	Amount   float64
	Currency string
}
