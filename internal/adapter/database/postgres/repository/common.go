package repository

import (
	"gorm.io/gorm"
	"log/slog"
	"orderbook/internal/core/model"
)

type CommonRepository struct {
	db *gorm.DB
}

func NewCommonRepository(db *gorm.DB) *CommonRepository {
	return &CommonRepository{
		db,
	}
}

func (cr *CommonRepository) GetTimeLimit(id string) (*model.TimeLimit, error) {
	var timeLimit *model.TimeLimit
	result := cr.db.Model(&model.TimeLimit{}).
		Select("*").
		Where("id = ?", id).
		First(&timeLimit)
	if result.Error != nil {
		slog.Info("GetTimeLimit error", result.Error.Error())
		return nil, result.Error
	}
	return timeLimit, nil
}
