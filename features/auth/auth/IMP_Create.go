package auth

import (
	"time"

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

func (S *Service) Create(app micro.IContext, task *CreateTask) (interface{}, error) {

	stype := dynamic.StringValue(task.Type, AuthType_JSON)
	expires := int64(task.Expires)
	key := task.Key

	if expires == 0 {
		expires = 60
	}

	var output interface{} = nil

	if stype == AuthType_JSON {
		err := json.Unmarshal([]byte(task.Value), &output)
		if err != nil {
			return nil, err
		}
	} else {
		output = task.Value
	}

	cli, prefix, err := app.GetRedis("default")

	if err != nil {
		return nil, err
	}

	_, err = cli.Set(prefix+key, task.Value, time.Second*time.Duration(expires)).Result()

	if err != nil {
		return nil, err
	}

	// MQ 消息
	app.SendMessage(task.GetName(), map[string]interface{}{"key": task.Key, "expires": task.Expires, "value": output})

	return output, nil
}
