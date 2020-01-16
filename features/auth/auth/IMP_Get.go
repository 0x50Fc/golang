package auth

import (
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Get(app micro.IContext, task *GetTask) (interface{}, error) {

	key := task.Key
	stype := dynamic.StringValue(task.Type, AuthType_JSON)

	cli, prefix, err := app.GetRedis("default")

	if err != nil {
		return nil, err
	}

	text, err := cli.Get(prefix + key).Result()

	if err == nil && text != "" {

		if stype == AuthType_JSON {
			var output interface{} = nil
			err = json.Unmarshal([]byte(text), &output)
			if err != nil {
				return nil, err
			}
			return output, nil
		}

		return text, nil
	}

	return nil, micro.NewError(ERROR_NOT_FOUND, "未找到验证对象")

}
