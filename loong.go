package loong

import (
	"github.com/panjf2000/loong/backends"
	"github.com/panjf2000/loong/brokers"
	"github.com/panjf2000/loong/config"
	"github.com/panjf2000/loong/factory"
)

type LoongClient struct {
	conf        config.Config
	concurrency int
	broker      brokers.StandardBroker
	backend     backends.StandardBackender
}

func NewLoongClient(concurrency int, conf config.Config) (*LoongClient, error) {
	broker, _ := factory.ProduceBroker(conf)
	backend, _ := factory.ProduceBackend(conf)
	cls := &LoongClient{
		conf:    conf,
		concurrency: concurrency,
		broker:  broker,
		backend: backend,
	}
	return cls, nil
}
