package backends

import "github.com/panjf2000/loong/protocol"

type StandardBackender interface{
	GetResult(taskID string) (protocol.ResultMessage, error)
}