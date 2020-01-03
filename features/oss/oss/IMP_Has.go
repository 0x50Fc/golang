package oss

import (
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Has(app micro.IContext, task *HasTask) (interface{}, error) {

	name := dynamic.StringValue(task.Name, "oss")

	source, err := GetSource(app, name)

	if err != nil {
		return nil, err
	}

	err = source.Has(task.Key)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
