package nsq

import (
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/mq"
	Q "github.com/nsqio/go-nsq"
)

func init() {
	mq.AddConsumer("nsq", func(config interface{}) (mq.Consumer, error) {
		return NewNSQConsumer(dynamic.StringValue(dynamic.Get(config, "addr"), ""),
			dynamic.StringValue(dynamic.Get(config, "topic"), ""),
			dynamic.StringValue(dynamic.Get(config, "channel"), ""))
	})
	mq.AddProducer("nsq", func(config interface{}) (mq.Producer, error) {
		return NewNSQProducer(dynamic.StringValue(dynamic.Get(config, "addr"), ""),
			dynamic.StringValue(dynamic.Get(config, "topic"), ""))
	})
}

type NSQConsumer struct {
	q    *Q.Consumer
	addr string
}

func NewNSQConsumer(addr string, topic string, channel string) (*NSQConsumer, error) {

	config := Q.NewConfig()
	q, err := Q.NewConsumer(topic, channel, config)

	if err != nil {
		return nil, err
	}

	return &NSQConsumer{q: q, addr: addr}, nil
}

type NSQConsumerHandler struct {
	cb mq.Listener
}

func (Q *NSQConsumerHandler) HandleMessage(message *Q.Message) error {

	var object interface{} = nil

	if message.Body != nil {
		_ = json.Unmarshal(message.Body, &object)
	}

	if object == nil {
		return nil
	}

	name := dynamic.StringValue(dynamic.Get(object, "name"), "")

	if name == "" {
		return nil
	}

	return Q.cb(name, dynamic.Get(object, "data"))
}

func (C *NSQConsumer) Close() {
	C.q.DisconnectFromNSQLookupd(C.addr)
	C.q.Stop()
}

func (C *NSQConsumer) Open(cb mq.Listener, concurrency int) error {
	C.q.AddConcurrentHandlers(&NSQConsumerHandler{cb: cb}, concurrency)
	return C.q.ConnectToNSQLookupd(C.addr)
}

type NSQProducer struct {
	q     *Q.Producer
	topic string
}

func NewNSQProducer(addr string, topic string) (*NSQProducer, error) {

	config := Q.NewConfig()
	q, err := Q.NewProducer(addr, config)

	if err != nil {
		return nil, err
	}

	return &NSQProducer{q: q, topic: topic}, nil
}

func (Q *NSQProducer) Close() {
	Q.q.Stop()
}

func (Q *NSQProducer) Send(name string, data interface{}) error {
	b, err := json.Marshal(map[string]interface{}{"name": name, "data": data})
	if err != nil {
		return err
	}
	return Q.q.Publish(Q.topic, b)
}
