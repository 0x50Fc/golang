package user

import (
	"fmt"
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

/**
 * JSON 数据叠加
 */
func Subjoin(v string, output interface{}) interface{} {
	var data interface{} = nil

	_ = json.Unmarshal([]byte(v), &data)

	if data == nil {
		data = map[string]interface{}{}
	}

	dynamic.Each(output, func(key interface{}, value interface{}) bool {
		dynamic.Set(data, dynamic.StringValue(key, ""), value)
		return true
	})

	return data
}

func (S *Service) InfoSet(app micro.IContext, task *InfoSetTask) (interface{}, error) {

	var output interface{} = nil

	if task.Type == InfoType_JSON && task.Value != nil {
		err := json.Unmarshal([]byte(dynamic.StringValue(task.Value, "")), &output)
		if err != nil {
			return nil, err
		}
	}

	conn, prefix, err := app.GetDB("wd")

	prefix = Prefix(app, prefix, task.Uid)

	if err != nil {
		return nil, err
	}

	v := Info{}

	p, err := db.Get(conn, &v, prefix, " WHERE `uid`=? AND `key`=?", task.Uid, task.Key)

	if err != nil {
		return nil, err
	}

	if task.Value != nil {
		if task.Type == InfoType_JSON {
			if v.Value != "" {
				output = Subjoin(v.Value, output)
				b, _ := json.Marshal(output)
				v.Value = string(b)
			} else {
				v.Value = dynamic.StringValue(task.Value, "")
			}
		} else {
			v.Value = dynamic.StringValue(task.Value, "")
			output = v.Value
		}

	}

	if p != nil {

		_, err = db.Update(conn, &v, prefix)

		if err != nil {
			return nil, err
		}

	} else {
		v.Uid = task.Uid
		v.Key = task.Key
		_, err = db.Insert(conn, &v, prefix)
		if err != nil {
			return nil, err
		}
	}

	{
		cache, _ := app.GetCache("default")

		if cache != nil {

			expires := time.Duration(dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"cache", "maxSecond"}), 1800)) * time.Second

			cache.Set(fmt.Sprintf("%d_%s", task.Uid, task.Key), v.Value, expires)

		}

	}

	// MQ 消息
	app.SendMessage(task.GetName(), &v)

	return output, nil
}
