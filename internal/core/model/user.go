package model

import (
	"gorm.io/gorm"
)

type UserType string

const (
	Free    UserType = "free"
	Premium UserType = "premium"
)

type User struct {
	gorm.Model
	Email       string   `gorm:"unique;not null" json:"email"`
	UserType    UserType `gorm:"default:free" json:"userType"`
	DisplayName *string  `json:"displayName"`
}
