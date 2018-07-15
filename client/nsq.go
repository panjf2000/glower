package common

import (
	"fmt"

	"github.com/nsqio/go-nsq"
)

type NSQClient struct {
	config      *nsq.Config
	consumer    *nsq.Consumer
	producer    *nsq.Producer
	nsqds       []string
	nsqlookupds []string
	concurrency int
	channel     string
	topic       string
	err         error
}

//初始化消费端
func NewClient(topic, channel string) NSQClient {
	return NSQClient{
		config:      nsq.NewConfig(),
		channel:     channel,
		topic:       topic,
		concurrency: 1,
	}
}

func (n *NSQClient) SetMap(options map[string]interface{}) {
	for k, v := range options {
		n.Set(k, v)
	}
}

func (n *NSQClient) Set(option string, value interface{}) {
	switch option {
	case "topic":
		n.topic = value.(string)
	case "channel":
		n.channel = value.(string)
	case "concurrency":
		n.concurrency = value.(int)
	case "nsqd":
		n.nsqds = []string{value.(string)}
	case "nsqlookupd":
		n.nsqlookupds = []string{value.(string)}
	case "nsqds":
		s, err := strings(value)
		if err != nil {
			n.err = fmt.Errorf("%q: %v", option, err)
			return
		}
		n.nsqds = s
	case "nsqlookupds":
		s, err := strings(value)
		if err != nil {
			n.err = fmt.Errorf("%q: %v", option, err)
			return
		}
		n.nsqlookupds = s
	default:
		err := n.config.Set(option, value)
		if err != nil {
			n.err = err
		}
	}
}

func (n *NSQClient) Receive(handler nsq.Handler) error {

	if n.err != nil {
		return n.err
	}

	client, err := nsq.NewConsumer(n.topic, n.channel, n.config)
	if err != nil {
		return err
	}
	n.consumer = client
	client.AddConcurrentHandlers(handler, n.concurrency)
	return n.connect()
}

//连接到nsqd
func (n *NSQClient) connect() error {

	if len(n.nsqds) == 0 && len(n.nsqlookupds) == 0 {
		return fmt.Errorf(`at least one "nsqd" or "nsqlookupd" address must be configured`)
	}

	if len(n.nsqds) > 0 {
		err := n.consumer.ConnectToNSQDs(n.nsqds)
		if err != nil {
			return err
		}
	}
	if len(n.nsqlookupds) > 0 {
		err := n.consumer.ConnectToNSQLookupds(n.nsqlookupds)
		if err != nil {
			return err
		}
	}
	return nil
}

//stop and wait
func (n *NSQClient) StopReceive() error {
	n.consumer.Stop()
	<-n.consumer.StopChan
	return nil
}

func strings(v interface{}) ([]string, error) {
	switch v.(type) {
	case []string:
		return v.([]string), nil
	case []interface{}:
		var ret []string
		for _, e := range v.([]interface{}) {
			s, ok := e.(string)
			if !ok {
				return nil, fmt.Errorf("string expected")
			}
			ret = append(ret, s)
		}
		return ret, nil
	default:
		return nil, fmt.Errorf("strings expected")
	}
}
