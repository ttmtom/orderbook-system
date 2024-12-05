package module

import (
	"gorm.io/gorm"
	"orderbook/internal/adapter/database/postgres/repository"
)

type CommonModule struct {
	Repository *repository.CommonRepository
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
