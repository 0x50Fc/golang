package oss

import (
	"time"

	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

type ISource interface {
	Get(key string) ([]byte, error)
	GetURL(key string) string
	GetSignURL(key string, expires time.Duration) (string, error)
	Put(key string, data []byte, header map[string]string) error
	PutSignURL(key string, expires time.Duration) (string, error)
	PostSignURL(key string, expires time.Duration) (string, map[string]string, error)
	Del(key string) error
	Has(key string) error
}

var sources = map[string](func(config interface{}) (ISource, error)){}

func AddSource(stype string, fn func(config interface{}) (ISource, error)) {
	sources[stype] = fn
}

func OpenSource(stype string, config interface{}) (ISource, error) {
	fn, ok := sources[stype]
	if ok {
		return fn(config)
	}
	return nil, micro.NewError(ERROR_NOT_FOUND, "未找到提供源")
}

func GetSource(app micro.IContext, name string) (ISource, error) {

	v, err := app.GetSharedObject(name, func() (micro.SharedObject, error) {

		config := dynamic.Get(app.GetConfig(), name)
		stype := dynamic.StringValue(dynamic.Get(config, "type"), "")

		v, err := OpenSource(stype, config)

		if err != nil {
			return nil, err
		}

		return v, nil
	})

	if err != nil {
		return nil, err
	}

	return v.(ISource), nil
}
