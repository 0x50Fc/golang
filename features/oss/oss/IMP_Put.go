package oss

import (
	"encoding/base64"
	"mime"
	"path/filepath"
	"time"

	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Put(app micro.IContext, task *PutTask) (interface{}, error) {

	name := dynamic.StringValue(task.Name, "oss")

	stype := dynamic.StringValue(task.Type, Type_URL)

	source, err := GetSource(app, name)

	if err != nil {
		return nil, err
	}

	ext := filepath.Ext(task.Key)
	contentType := mime.TypeByExtension(ext)

	header := map[string]string{}

	if contentType != "" {
		header["Content-Type"] = contentType
	}

	switch stype {
	case Type_Text:
		{
			err = source.Put(task.Key, []byte(dynamic.StringValue(task.Content, "")), header)
			if err != nil {
				return nil, err
			}
			return nil, nil
		}
	case Type_Base64:
		{
			b, err := base64.StdEncoding.DecodeString(dynamic.StringValue(task.Content, ""))
			if err != nil {
				return nil, err
			}
			err = source.Put(task.Key, b, header)
			if err != nil {
				return nil, err
			}
			return nil, nil
		}
	default:
		{
			v, err := source.PutSignURL(task.Key, time.Second*time.Duration(dynamic.IntValue(task.Expires, 60)))

			if err != nil {
				return nil, err
			}

			return v, nil
		}
	}

}
