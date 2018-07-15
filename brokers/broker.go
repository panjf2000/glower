package brokers

import "github.com/panjf2000/loong/protocol"

type StandardBroker interface {
	Publish(proto *protocol.Protocol) error
	Receive() (protocol.ProtoMessage, error)
}
