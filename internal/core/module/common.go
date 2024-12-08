package module

import (
	"gorm.io/gorm"
	"orderbook/internal/adapter/database/postgres/repository"
	"orderbook/internal/core/port"
)

type CommonModule struct {
	Repository port.CommonRepository
	//Service    *service.CommonService
}

func NewCommonModule(connection *gorm.DB) *CommonModule {
	resp := repository.NewCommonRepository(connection)
	//svc := service.NewCommonService(resp)

	return &CommonModule{
		Repository: resp,
		//Service:    svc,
	}
}
