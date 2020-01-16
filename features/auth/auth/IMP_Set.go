package auth

import (
	"time"

	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Set(app micro.IContext, task *SetTask) (interface{}, error) {

	stype := dynamic.StringValue(task.Type, AuthType_JSON)
	expires := time.Duration(dynamic.IntValue(task.Expires, 0))
	key := task.Key

	var output interface{} = nil

	if stype == AuthType_JSON && task.Value != nil {
		err := json.Unmarshal([]byte(dynamic.StringValue(task.Value, "")), &output)
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

	if stype == AuthType_JSON {
		text, err := cli.Get(prefix + key).Result()
		if err == nil && text != "" {
			output = Subjoin(text, output)
		}
	}

	if stype == AuthType_JSON {
		b, _ := json.Marshal(output)
		_, err = cli.Set(prefix+key, string(b), time.Second*expires).Result()
	} else {
		_, err = cli.Set(prefix+key, task.Value, time.Second*expires).Result()
	}

	if err != nil {
		return nil, err
	}

	// MQ 消息
	app.SendMessage(task.GetName(), map[string]interface{}{"key": task.Key, "expires": task.Expires, "value": output})

	return output, nil
}
