package ali

import (
	"log"
	"time"

	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/mq"
	Ali "github.com/aliyun/aliyun-mns-go-sdk"
)

func init() {
	mq.AddConsumer("ali", func(config interface{}) (mq.Consumer, error) {
		return NewAliConsumer(dynamic.StringValue(dynamic.Get(config, "url"), ""),
			dynamic.StringValue(dynamic.Get(config, "accessKey"), ""),
			dynamic.StringValue(dynamic.Get(config, "secretKey"), ""),
			dynamic.StringValue(dynamic.Get(config, "queue"), ""))
	})
	mq.AddProducer("ali", func(config interface{}) (mq.Producer, error) {
		return NewAliProducer(dynamic.StringValue(dynamic.Get(config, "url"), ""),
			dynamic.StringValue(dynamic.Get(config, "accessKey"), ""),
			dynamic.StringValue(dynamic.Get(config, "secretKey"), ""),
			dynamic.StringValue(dynamic.Get(config, "topic"), ""))
	})
}

type AliConsumer struct {
	client Ali.MNSClient
	queue  Ali.AliMNSQueue
	closed bool
}

func NewAliConsumer(url string, accessKey string, secretKey string, queue string) (*AliConsumer, error) {

	cli := Ali.NewAliMNSClient(url, accessKey, secretKey)

	q := Ali.NewMNSQueue(queue, cli)

	return &AliConsumer{queue: q, client: cli, closed: false}, nil
}

func (C *AliConsumer) Close() {
	if C.closed {
		return
	}
	C.closed = true
}

func (C *AliConsumer) Open(cb mq.Listener, concurrency int) error {

	if concurrency < 1 {
		concurrency = 1
	}

	respChan := make(chan Ali.MessageReceiveResponse, concurrency)
	errChan := make(chan error, concurrency)

	for i := 0; i < concurrency; i++ {

		go func() {
			for !C.closed {
				select {
				case resp := <-respChan:
					{
						{
							var body interface{} = nil
							err := json.Unmarshal([]byte(resp.MessageBody), &body)
							name := dynamic.StringValue(dynamic.Get(body, "name"), "")
							if name != "" {
								err = cb(name, dynamic.Get(body, "data"))
								if err != nil {
									log.Println(err)
									time.Sleep(6 * time.Second)
									continue
								}
							}
						}
						if ret, e := C.queue.ChangeMessageVisibility(resp.ReceiptHandle, 5); e != nil {
							log.Println(e)
						} else {
							if e := C.queue.DeleteMessage(ret.ReceiptHandle); e != nil {
								log.Println(e)
							}
						}
					}
				case err := <-errChan:
					{
						log.Println(err)
					}
				}
			}
		}()
	}

	go func() {
		for !C.closed {
			C.queue.ReceiveMessage(respChan, errChan, 30)
		}
	}()

	return nil
}

type AliProducer struct {
	client Ali.MNSClient
	topic  Ali.AliMNSTopic
}

func NewAliProducer(url string, accessKey string, secretKey string, topic string) (*AliProducer, error) {

	cli := Ali.NewAliMNSClient(url, accessKey, secretKey)

	tp := Ali.NewMNSTopic(topic, cli)

	return &AliProducer{topic: tp, client: cli}, nil
}

func (Q *AliProducer) Close() {

}

func (Q *AliProducer) Send(name string, data interface{}) error {
	b, err := json.Marshal(map[string]interface{}{"name": name, "data": data})
	if err != nil {
		return err
	}
	msg := Ali.MessagePublishRequest{
		MessageBody: string(b),
	}
	_, err = Q.topic.PublishMessage(msg)
	return err
}
