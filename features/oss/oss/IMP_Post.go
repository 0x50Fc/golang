package oss

import (
	"time"

	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Post(app micro.IContext, task *PostTask) (*PostData, error) {

	name := dynamic.StringValue(task.Name, "oss")

	source, err := GetSource(app, name)

	if err != nil {
		return nil, err
	}

	v, data, err := source.PostSignURL(task.Key, time.Second*time.Duration(dynamic.IntValue(task.Expires, 60)))

	if err != nil {
		return nil, err
	}

	return &PostData{Url: v, Data: data}, nil

}
