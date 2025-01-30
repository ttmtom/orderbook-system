package repository

import (
	"gorm.io/gorm"
	"orderbook/internal/core/model"
	"orderbook/internal/core/port"
)

type AdminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(connection *gorm.DB) port.AdminRepository {
	return &AdminRepository{
		connection,
	}
}

func (a AdminRepository) GetAdmin(email string) (*model.AdminUser, error) {
	//TODO implement me
	panic("implement me")
}
