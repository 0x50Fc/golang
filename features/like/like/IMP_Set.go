package like

import (
	"fmt"
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Set(app micro.IContext, task *SetTask) (*Like, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	prefix = Prefix(app, prefix, task.Tid)

	like := Like{}

	iid := dynamic.IntValue(task.Iid, 0)

	isAdd := false

	err = db.Transaction(conn, func(conn db.Database) error {

		rs, err := db.Query(conn, &like, prefix, " WHERE tid = ? AND uid = ? AND iid=? ", task.Tid, task.Uid, iid)

		if err != nil {
			return err
		}

		if rs.Next() {
			scaner := db.NewScaner(&like)
			err = scaner.Scan(rs)
		}

		rs.Close()

		if err != nil {
			return err
		}

		if like.Id == 0 {
			like.Tid = task.Tid
			like.Iid = iid
			like.Uid = task.Uid
			like.Ctime = time.Now().Unix()
		}

		keys := map[string]bool{}

		if task.Options != nil {

			options := map[string]interface{}{}

			dynamic.Each(like.Options, func(key interface{}, value interface{}) bool {
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

			like.Options = options
			keys["options"] = true
		}

		if like.Id == 0 {
			isAdd = true
			_, err = db.Insert(conn, &like, prefix)
		} else if len(keys) > 0 {
			_, err = db.UpdateWithKeys(conn, &like, prefix, keys)
		}

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	//清除缓存
	{
		cache, err := app.GetCache("default")

		if err == nil {

			cache.Del(fmt.Sprintf("%d", task.Tid))

		}
	}

	if isAdd {
		// MQ 消息
		app.SendMessage(task.GetName(), &like)
	}

	return &like, nil
}
