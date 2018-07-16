package brokers

import (
	"github.com/panjf2000/loong/protocol"
	"github.com/panjf2000/loong/config"
)

type StandardBroker interface {
	Publish(proto *protocol.Protocol) error
	Receive() (protocol.ProtoMessage, error)
}

func NewNsqBroker(config config.BrokerConfig) () {
	
}