package model

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID           uint         `gorm:"primarykey" json:"id"`
	IDHash       string       `gorm:"unique;not null" json:"idHash"`
	Email        string       `gorm:"unique;not null" json:"email"`
	PasswordHash string       `gorm:"not null" json:"passwordHash"`
	DisplayName  *string      `json:"displayName"`
	CreatedAt    time.Time    `json:"createdAt"`
	UpdatedAt    time.Time    `json:"updatedAt"`
	DeletedAt    sql.NullTime `gorm:"index" json:"deletedAt"`
}
