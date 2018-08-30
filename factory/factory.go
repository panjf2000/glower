package factory

import (
	"github.com/panjf2000/loong/brokers"
	"github.com/panjf2000/loong/config"
)

type BrokerFactory struct {
	brokerConf config.BrokerConfig
	backendConf config.BrokerConfig
}

func (bf *BrokerFactory) ProduceBroker() (brokers.StandardBroker, error) {
	return nil, nil
}
