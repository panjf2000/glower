package factory

import (
	"github.com/panjf2000/loong/brokers"
	"github.com/panjf2000/loong/backends"
	"github.com/panjf2000/loong/config"
)

func ProduceBroker(conf config.Config) (brokers.StandardBroker, error) {
	return nil, nil
}

func ProduceBackend(conf config.Config) (backends.standardBackend, error) {
	return nil, nil
}
