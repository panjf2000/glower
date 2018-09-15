package brokers

import (
	"github.com/panjf2000/loong/protocol"
)

type StandardBroker interface {
	Publish(message protocol.ProtoMessage) error
	Receive() (protocol.ProtoMessage, error)
}
