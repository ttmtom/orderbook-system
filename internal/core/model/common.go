package model

type TimeLimit struct {
	ID   string `gorm:"primarykey" json:"id"`
	Time uint   `json:"time"`
}
