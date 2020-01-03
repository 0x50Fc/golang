package oss

import (
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Rm(app micro.IContext, task *RmTask) (interface{}, error) {

	name := dynamic.StringValue(task.Name, "oss")

	source, err := GetSource(app, name)

	if err != nil {
		return nil, err
	}

	err = source.Del(task.Key)

	if err != nil {
		return nil, err
	}

	return nil, nil

}
