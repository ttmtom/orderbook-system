package port

type ConsumerGroup interface {
	StartPolling()
	StopPolling()
}
