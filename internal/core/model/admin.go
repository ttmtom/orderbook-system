package model

type AdminUser struct {
	UserBase
	Role string `gorm:"not null" json:"role"`
}
