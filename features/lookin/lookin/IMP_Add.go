package lookin

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Add(app micro.IContext, task *AddTask) (*AddData, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	prefix = Prefix(app, prefix, task.Tid)

	iid := dynamic.IntValue(task.Iid, 0)

	items := []*Lookin{&Lookin{Fuid: task.Uid, Flevel: 0}}

	fcode := dynamic.StringValue(task.Fcode, "")

	if task.Fcode != nil {
		ids, err := DeocdeString(fcode)
		if err != nil {
			return nil, err
		}
		n := len(ids)
		for i, id := range ids {
			if id != 0 && id != task.Uid {
				items = append(items, &Lookin{Fuid: dynamic.IntValue(task.Fuid, 0), Flevel: int32(n - i)})
			}
		}
	} else if task.Fuid != nil {
		items = append(items, &Lookin{Fuid: dynamic.IntValue(task.Fuid, 0), Flevel: 1})
	}

	v := Lookin{}

	for _, item := range items {
		p, err := db.Get(conn, &v, prefix, " WHERE tid=? AND iid=? AND uid=? AND fuid=? ORDER BY flevel DESC LIMIT 1", task.Tid, iid, task.Uid, item.Fuid)
		if err != nil {
			return nil, err
		}
		if p == nil {
			item.Tid = task.Tid
			item.Iid = iid
			item.Uid = task.Uid
			item.Fcode = fcode
			if task.Options != nil {
				json.Unmarshal([]byte(dynamic.StringValue(task.Options, "")), &task.Options)
			}
			item.Ctime = time.Now().Unix()
			_, err = db.Insert(conn, item, prefix)
			if err != nil {
				return nil, err
			}
		} else if v.Flevel < item.Flevel {
			*item = v
		} else {

			v.Flevel = item.Flevel
			v.Fuid = item.Fuid
			v.Fcode = fcode

			*item = v

			keys := map[string]bool{"flevel": true, "fuid": true, "fcode": true}

			if task.Options != nil {
				var options interface{} = nil
				json.Unmarshal([]byte(dynamic.StringValue(task.Options, "")), &options)
				item.Options = db.Merge(v.Options, options)
				v.Options = options
				keys["options"] = true
			}

			_, err = db.UpdateWithKeys(conn, &v, prefix, keys)

			if err != nil {
				return nil, err
			}

		}
	}

	//清除缓存
	{
		cache, err := app.GetCache("default")

		if err == nil {

			cache.Del(fmt.Sprintf("%d_%d", task.Tid, iid))

		}
	}

	data := AddData{Items: items}

	app.SendMessage(task.GetName(), &data)

	return &data, nil
}
