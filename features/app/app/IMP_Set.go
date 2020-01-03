package app

import (
	"fmt"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Set(app micro.IContext, task *SetTask) (*App, error) {
	v := App{}

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	prefix = Prefix(app, prefix, task.Id)

	p, err := db.Get(conn, &v, prefix, "WHERE id=?", task.Id)

	if err != nil {
		return nil, err
	}

	if p == nil {
		return nil, micro.NewError(ERROR_NOT_FOUND, "未找到App")
	}

	keys := map[string]bool{}

	if task.Uid != nil {

		v.Uid = dynamic.IntValue(task.Uid, 0)

		keys["uid"] = true
	}

	if task.Title != nil {

		v.Title = dynamic.StringValue(task.Title, "")

		keys["title"] = true
	}

	if task.Options != nil {

		options := map[string]interface{}{}

		dynamic.Each(v.Options, func(key interface{}, value interface{}) bool {
			options[dynamic.StringValue(key, "")] = value
			return true
		})

		text := dynamic.StringValue(task.Options, "")

		var data interface{} = nil

		json.Unmarshal([]byte(text), &data)

		dynamic.Each(data, func(key interface{}, value interface{}) bool {
			options[dynamic.StringValue(key, "")] = value
			return true
		})

		v.Options = options
		keys["options"] = true
	}

	if len(keys) > 0 {
		_, err = db.UpdateWithKeys(conn, &v, prefix, keys)
		if err != nil {
			return nil, err
		}
	}

	cache, _ := app.GetCache("default")

	if cache != nil {
		cache.Del(fmt.Sprintf("%d", task.Id))
	}

	return &v, nil
}
