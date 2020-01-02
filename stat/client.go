package stat

import (
	"errors"
	"time"
)

type Client interface {
	Close()
	Write(name string, tags map[string]string, fields map[string]interface{}, tv time.Time) error
}

type Object struct {
	Name   string                 `json:"name"`
	Tags   map[string]string      `json:"tags"`
	Fields map[string]interface{} `json:"fields"`
	Tv     int64                  `json:"tv"`
}

var openlibs = map[string](func(config interface{}) (Client, error)){}

func AddOpenlib(stype string, fn func(config interface{}) (Client, error)) {
	openlibs[stype] = fn
}

func OpenClient(stype string, config interface{}) (Client, error) {
	fn, ok := openlibs[stype]
	if ok {
		return fn(config)
	}
	return nil, errors.New("未找到实现类")
}
