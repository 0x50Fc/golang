package oss

import (
	"encoding/base64"
	"time"

	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Get(app micro.IContext, task *GetTask) (interface{}, error) {

	name := dynamic.StringValue(task.Name, "oss")

	stype := dynamic.StringValue(task.Type, Type_URL)

	source, err := GetSource(app, name)

	if err != nil {
		return nil, err
	}

	header := map[string]string{}

	if task.Header != nil {
		var data interface{} = nil
		json.Unmarshal([]byte(dynamic.StringValue(task.Header, "")), &data)
		dynamic.Each(data, func(key interface{}, value interface{}) bool {
			header[dynamic.StringValue(key, "")] = dynamic.StringValue(value, "")
			return true
		})
	}

	switch stype {
	case Type_Text:
		{
			b, err := source.Get(task.Key, header)
			if err != nil {
				return nil, err
			}
			return string(b), nil
		}
	case Type_Base64:
		{
			b, err := source.Get(task.Key, header)
			if err != nil {
				return nil, err
			}
			return base64.StdEncoding.EncodeToString(b), nil
		}
	default:
		if task.Expires == nil {

			return source.GetURL(task.Key), nil

		} else {

			u, err := source.GetSignURL(task.Key, time.Second*time.Duration(dynamic.IntValue(task.Expires, 0)), header)

			if err != nil {
				return nil, err
			}

			return u, nil
		}
	}

}
