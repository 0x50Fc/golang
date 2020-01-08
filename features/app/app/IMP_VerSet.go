package app

import (
	"fmt"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) VerSet(app micro.IContext, task *VerSetTask) (*Ver, error) {
	v := Ver{}

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	prefix = Prefix(app, prefix, task.Appid)

	p, err := db.Get(conn, &v, prefix, "WHERE appid=? AND ver=?", task.Appid, task.Ver)

	if err != nil {
		return nil, err
	}

	if p == nil {
		return nil, micro.NewError(ERROR_NOT_FOUND, "未找到App版本")
	}

	keys := map[string]bool{}

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

	if task.Info != nil {

		info := map[string]interface{}{}

		dynamic.Each(v.Info, func(key interface{}, value interface{}) bool {
			info[dynamic.StringValue(key, "")] = value
			return true
		})

		text := dynamic.StringValue(task.Info, "")

		var data interface{} = nil

		json.Unmarshal([]byte(text), &data)

		dynamic.Each(data, func(key interface{}, value interface{}) bool {
			info[dynamic.StringValue(key, "")] = value
			return true
		})

		v.Info = info
		keys["info"] = true
	}

	if len(keys) > 0 {
		_, err = db.UpdateWithKeys(conn, &v, prefix, keys)
		if err != nil {
			return nil, err
		}
	}

	cache, _ := app.GetCache("default")

	if cache != nil {
		cache.DelItem(fmt.Sprintf("%d", task.Appid), fmt.Sprintf("%d", task.Ver))
	}

	return &v, nil
}
