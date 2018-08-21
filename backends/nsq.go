package backends 
import (
	"github.com/nsqio/go-nsq"
)
type NsqBackend struct{
	nsqd string
	config nsq.Config
	consumer *nsq.Consumer
	channel string
}