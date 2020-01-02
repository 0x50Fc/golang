package user

import (
	"fmt"
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) InfoGet(app micro.IContext, task *InfoGetTask) (interface{}, error) {

	expires := time.Duration(dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"cache", "maxSecond"}), 1800)) * time.Second

	var data interface{} = nil

	cache, _ := app.GetCache("default")

	if cache != nil {

		text, err := cache.Get(fmt.Sprintf("%d_%s", task.Uid, task.Key), expires)

		if err == nil {
			if text == "__nil" {
				return nil, nil
			}
			if task.Type == InfoType_JSON {
				err = json.Unmarshal([]byte(dynamic.StringValue(text, "")), &data)
				if err == nil {
					return data, nil
				}
			} else {
				data = dynamic.StringValue(text, "")
				return data, nil
			}
		}
	}

	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}

	prefix = Prefix(app, prefix, task.Uid)

	v := Info{}

	p, err := db.Get(conn, &v, prefix, " WHERE `uid`=? AND `key`=?", task.Uid, task.Key)

	if p == nil {

		if cache != nil {
			cache.Set(fmt.Sprintf("%d_%s", task.Uid, task.Key), "__nil", expires)
		}

		return nil, nil
	}

	if task.Type == InfoType_JSON {

		err = json.Unmarshal([]byte(v.Value), &data)

		if err != nil {
			return nil, err
		}

	} else {
		data = v.Value
	}

	if cache != nil {
		cache.Set(fmt.Sprintf("%d_%s", task.Uid, task.Key), v.Value, expires)
	}

	return data, nil

}
