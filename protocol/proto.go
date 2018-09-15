package protocol

import "time"

type Params struct {
	Type  string
	Value interface{}
}

type Protocol map[string]interface{}

type ProtoMessage struct {
	UUID       string
	Name       string
	RoutingKey string
	ETA        *time.Time
	Headers    Protocol
	params     []Params
}

type ResultMessage struct {
}
