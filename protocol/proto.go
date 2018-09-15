package protocol

import "time"

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
}

type ResultMessage struct {
	results []Params
}
