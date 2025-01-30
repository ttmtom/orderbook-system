package service

import (
	"orderbook/internal/core/model"
	"orderbook/internal/core/port"
)

type AdminService struct {
}

func NewAdminService() port.AdminService {
	return &AdminService{}
}

func (a AdminService) AdminLogin(email, password string) (*model.AdminUser, string, error) {
	//TODO implement me
	panic("implement me")
}
