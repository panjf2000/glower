package loong

import "github.com/panjf2000/loong/protocol"

type LoongWorker struct {
	client *LoongClient
}

func (w *LoongWorker) Process(msg *protocol.ProtoMessage) error {
	return nil
}

func (w *LoongWorker) StartWorking() {
	for i := 0; i < w.client.concurrency; i++ {
		go func(workerID int) {
			for {
				w.client.broker.Processing(w)
			}
		}(i)
	}
}