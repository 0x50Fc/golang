package app

import (
	"fmt"
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) VerAdd(app micro.IContext, task *VerAddTask) (*Ver, error) {

	a := App{}
	v := Ver{}

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	prefix = Prefix(app, prefix, task.Appid)

	err = db.Transaction(conn, func(conn db.Database) error {

		p, err := db.Get(conn, &a, prefix, "WHERE id=?", task.Appid)

		if err != nil {
			return err
		}

		if p == nil {
			return micro.NewError(ERROR_NOT_FOUND, "未找到App")
		}

		v.Appid = task.Appid
		v.Ver = a.LastVer + 1

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
		}

		v.Ctime = time.Now().Unix()

		a.LastVer = v.Ver

		_, err = db.UpdateWithKeys(conn, &a, prefix, map[string]bool{"lastver": true})

		if err != nil {
			return err
		}

		_, err = db.Insert(conn, &v, prefix)

		if err != nil {
			return err
		}

		return nil
	})

	cache, _ := app.GetCache("default")

	if cache != nil {
		cache.Del(fmt.Sprintf("%d", task.Appid))
	}

	return &v, nil
}
