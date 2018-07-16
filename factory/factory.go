package factory

import (
	"github.com/panjf2000/loong/brokers"
)

type BrokerFactory struct {
}

func (bf BrokerFactory) ProduceBroker() (brokers.StandardBroker, error) {
	return nil, nil
}
