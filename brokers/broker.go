package brokers

type StandardBroker interface {
	Publish() error
	Receive()
}
