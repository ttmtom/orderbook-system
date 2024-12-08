package port

type EventRepository interface {
	PublishEvent(topic string, event any) error

	SetUpGroupConsumer(
		group string,
		topicMap map[string]func(event []byte) error,
		pollingInterval int,
		retry int,
	) ConsumerGroup
	CloseAll()
	StartPolling()
}
