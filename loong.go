package loong

import (
	"github.com/panjf2000/loong/backends"
	"github.com/panjf2000/loong/brokers"
)

type LoongClient struct {
	broker  brokers.StandardBroker
	backend backends.StandardBackender
}

func NewLoongClient(concurrency int, broker brokers.StandardBroker, backend backends.StandardBackender) {

}
