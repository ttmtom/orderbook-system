package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID           uint         `gorm:"primarykey" json:"id"`
	Email        string       `gorm:"unique;not null" json:"email"`
	PasswordHash string       `gorm:"not null" json:"passwordHash"`
	CreatedAt    time.Time    `json:"createdAt"`
	UpdatedAt    time.Time    `json:"updatedAt"`
	DeletedAt    sql.NullTime `gorm:"index" json:"deletedAt"`
	IDHash       string       `gorm:"unique;not null" json:"idHash"`
	DisplayName  *string      `json:"displayName"`
	LastLoginAt  time.Time    `json:"lastLoginAt"`
	LastAccessAt time.Time    `json:"lastAccessAt"`
}
