package config

type BrokerConfig struct {
	Class  string
	Broker string
	Queue  string
}

type BackendConfig struct {
	Class   string
	Backend string
	Queue   string
}
