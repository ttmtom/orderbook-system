package model

import "time"

type InstrumentType string

const (
	Withdrawal InstrumentType = "withdrawal"
	Deposit    InstrumentType = "deposit"
)

type Instrument struct {
	ID        uint `gorm:"primarykey" json:"id"`
	Wallet    Wallet
	Type      InstrumentType `gorm:"not null" json:"type"`
	Amount    float64        `gorm:"not null" json:"amount"`
	CreatedAt time.Time      `json:"createdAt"`
}
