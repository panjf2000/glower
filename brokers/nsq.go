package brokers
import (
	"github.com/nsqio/go-nsq"
)
type NsqBroker struct{
	nsqd string
	config nsq.Config
	producer *nsq.Producer	
	consumer *nsq.Consumer
}