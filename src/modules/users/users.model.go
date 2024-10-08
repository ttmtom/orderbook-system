package users

type User struct {
	ID       int
	Accounts []Account
}

type Account struct {
	ID     int
	Amount float64
	Type   string
}
