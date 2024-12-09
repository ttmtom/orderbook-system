package model

import "time"

type CryptoCurrency string

const (
	BTC CryptoCurrency = "BTC"
	ETH CryptoCurrency = "ETH"
)

type Wallet struct {
	ID       uint           `gorm:"primarykey" json:"id"`
	UserID   uint           `gorm:"foreignKey:User;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User     User           `gorm:"foreignkey:UserID"`
	Currency CryptoCurrency `gorm:"unique;not null" json:"currency"`
	Balance  float64        `gorm:"not null" json:"balance"`
	Locked   bool           `gorm:"not null" json:"locked"`
	Hold     float64        `gorm:"not null" json:"hold"`

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
	ID         uint       `gorm:"primarykey" json:"id"`
	Wallet     Wallet     `gorm:"foreignkey:WalletId"`
	Instrument Instrument `gorm:"foreignkey:InstrumentId"`
	Status     Status     `gorm:"not null" json:"status"`

	CreatedAt time.Time `json:"createdAt"`
}

type InstrumentType string

const (
	Withdrawal InstrumentType = "withdrawal"
	Deposit    InstrumentType = "deposit"
)

type Instrument struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Wallet    Wallet         `gorm:"foreignkey:WalletId"`
	User      User           `gorm:"foreignKey:UserId"`
	Type      InstrumentType `gorm:"not null" json:"type"`
	Amount    float64        `gorm:"not null" json:"amount"`
	CreatedAt time.Time      `json:"createdAt"`
}
