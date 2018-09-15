package protocol

import (
	"time"

	"github.com/panjf2000/loong/config"
)

type Params struct {
	Type  string
	Value interface{}
}

type Processor interface {
	Process(*ProtoMessage) error
}

type Protocol map[string]interface{}

type ProtoMessage struct {
	UUID       string
	Name       string
	RoutingKey string
	ETA        *time.Time
	Headers    Protocol
	Args       []Params
	conf       config.BrokerConfig
}

type ResultMessage struct {
	results []Params
}
