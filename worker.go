package loong

import (
	"github.com/panjf2000/loong/brokers"
	"github.com/panjf2000/loong/backends"
)

type Worker struct{
	broker brokers.StandardBroker
	backend backends.StandardBackender
}