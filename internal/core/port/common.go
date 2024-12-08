package port

import "orderbook/internal/core/model"

type CommonRepository interface {
	GetTimeLimit(id string) (*model.TimeLimit, error)
}
