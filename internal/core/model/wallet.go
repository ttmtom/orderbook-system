package model

import "time"

type CryptoCurrency string

const (
	BTC CryptoCurrency = "BTC"
	ETH CryptoCurrency = "ETH"
)

type Wallet struct {
	ID     uint `gorm:"primarykey" json:"id"`
	UserID uint `gorm:"foreignKey:User;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	//User     User           `gorm:"foreignkey:UserID"`
	Currency CryptoCurrency `gorm:"unique;not null" json:"currency"`
	Balance  float64        `gorm:"not null" json:"balance"`
	Blocked  bool           `gorm:"not null" json:"blocked"`
	Locked   float64        `gorm:"not null" json:"locked"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type TransactionType string

const (
	Deposit    TransactionType = "deposit"
	Withdrawal TransactionType = "withdrawal"
	Transfer   TransactionType = "transfer"
)

type Transaction struct {
	ID          uint            `gorm:"primarykey" json:"id"`
	IDHash      string          `gorm:"unique;not null" json:"idHash"`
	Amount      float64         `gorm:"not null" json:"balance"`
	Type        TransactionType `gorm:"not null" json:"type"`
	FromID      *uint           `gorm:"foreignKey:Wallet;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ToID        *uint           `gorm:"foreignKey:Wallet;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Description string          `json:"description"`
	CreatedAt   time.Time       `json:"createdAt"`
}

type TransactionEventType string

const (
	Pending   TransactionEventType = "pending"
	Success   TransactionType      = "success"
	Rejected  TransactionEventType = "rejected"
	Cancelled TransactionEventType = "cancelled"
)

type TransactionEvent struct {
	ID            uint                 `gorm:"primarykey" json:"id"`
	TransactionID uint                 `gorm:"foreignKey:Transaction;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Type          TransactionEventType `gorm:"not null" json:"type"`
	CreatedAt     time.Time            `json:"createdAt"`
}

type WalletActionType string

const (
	Increase WalletActionType = "increase"
	Decrease WalletActionType = "decrease"
	Lock     WalletActionType = "lock"
	Release  WalletActionType = "release"
	Commit   WalletActionType = "commit"
)

type WalletHistory struct {
	ID       uint `gorm:"primarykey" json:"id"`
	WalletID uint `gorm:"foreignKey:Wallet;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	//Wallet     Wallet     `gorm:"foreignkey:WalletId"`
	Action WalletActionType `gorm:"not null" json:"action"`

	CreatedAt time.Time `json:"createdAt"`
}
