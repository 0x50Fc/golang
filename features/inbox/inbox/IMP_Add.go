package inbox

import (
	"fmt"
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Add(app micro.IContext, task *AddTask) (*Inbox, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	iid := dynamic.IntValue(task.Iid, 0)
	prefix = Prefix(app, prefix, task.Uid)

	v := Inbox{}

	err = db.Transaction(conn, func(conn db.Database) error {

		_, err := db.Get(conn, &v, prefix, " WHERE mid=? AND uid=? AND iid=?", task.Mid, task.Uid, iid)

		if err != nil {
			return err
		}

		v.Uid = task.Uid
		v.Fuid = task.Fuid
		v.Type = task.Type | v.Type
		v.Mid = task.Mid
		v.Iid = iid
		v.Ctime = time.Now().Unix()

		if task.Ctime != nil {
			v.Ctime = dynamic.IntValue(task.Ctime, v.Ctime)
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

		if v.Id == 0 {
			_, err = db.Insert(conn, &v, prefix)

			if err != nil {
				return err
			}

		} else {

			_, err = db.Update(conn, &v, prefix)

			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	{
		//清除缓存
		cache, err := app.GetCache("default")
		if err == nil {
			cache.Del(fmt.Sprintf("%d", task.Uid))
		}
	}

	// MQ 消息
	app.SendMessage(task.GetName(), &v)

	return &v, nil
}
