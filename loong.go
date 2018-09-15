package loong

import (
	"github.com/panjf2000/loong/backends"
	"github.com/panjf2000/loong/brokers"
	"github.com/panjf2000/loong/config"
)

type LoongClient struct {
	conf config.Config
	concurrency int
	broker  brokers.StandardBroker
	backend backends.StandardBackender
}

func NewLoongClient(concurrency int, conf config.Config) (*LoongClient, error){
	cls := &LoongClient{
		conf: conf,
	}
	return cls, nil
}
