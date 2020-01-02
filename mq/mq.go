package mq

import (
	"errors"
)

type Listener func(name string, data interface{}) error

type Consumer interface {
	Close()
	Open(cb Listener, concurrency int) error
}

type Producer interface {
	Close()
	Send(name string, data interface{}) error
}

var consumers = map[string](func(config interface{}) (Consumer, error)){}
var producers = map[string](func(config interface{}) (Producer, error)){}

func AddConsumer(stype string, openlib func(config interface{}) (Consumer, error)) {
	consumers[stype] = openlib
}

func AddProducer(stype string, openlib func(config interface{}) (Producer, error)) {
	producers[stype] = openlib
}

func OpenConsumer(stype string, config interface{}) (Consumer, error) {
	v, ok := consumers[stype]
	if ok {
		return v(config)
	}
	return nil, errors.New("未找到 MQ 订阅实现")
}

func OpenProducer(stype string, config interface{}) (Producer, error) {
	v, ok := producers[stype]
	if ok {
		return v(config)
	}
	return nil, errors.New("未找到 MQ 发布实现")
}
