package model

import "time"

type CryptoCurrency string

const (
	BTC CryptoCurrency = "btc"
	ETH CryptoCurrency = "eth"
)

type Wallet struct {
	ID       uint `gorm:"primarykey" json:"id"`
	User     User
	Currency CryptoCurrency `gorm:"unique;not null" json:"currency"`
	Balance  float64        `gorm:"not null" json:"balance"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Status string

const (
	Pending   Status = "pending"
	Completed Status = "completed"
	Canceled  Status = "canceled"
	Failed    Status = "failed"
)

type WalletHistory struct {
	ID uint `gorm:"primarykey" json:"id"`
}
