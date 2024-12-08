package module

import (
	"orderbook/config"
	"orderbook/internal/adapter/kafka"
	"orderbook/internal/core/port"
)

type EventModule struct {
	Repository port.EventRepository
}

func NewEventModule(config *config.Config) *EventModule {
	eventManager := kafka.New(config)

	return &EventModule{eventManager}
}

func (em *EventModule) StartPolling() {
	em.Repository.StartPolling()
}

func (em *EventModule) CloseAll() {
	em.Repository.CloseAll()
}
