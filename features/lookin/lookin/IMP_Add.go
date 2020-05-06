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

	maxLevel := int(dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"lookin", "maxLevel"}), 3))

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	prefix = Prefix(app, prefix, task.Tid)

	iid := dynamic.IntValue(task.Iid, 0)

	items := []*Lookin{&Lookin{Fuid: task.Uid, Flevel: 0}}

	var iids []int64

	fcode := dynamic.StringValue(task.Fcode, "")

	if task.Fcode == nil {
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
		iids = append(ids, task.Uid)
		n += 1
		if n > maxLevel {
			iids = iids[n-maxLevel:]
		}
	} else if task.Fuid == nil {
		items = append(items, &Lookin{Fuid: dynamic.IntValue(task.Fuid, 0), Flevel: 1})
		iids = []int64{task.Uid}
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
			item.Options = task.Options
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

			keys := map[string]bool{"flevel": true, "fuid": true, "fcode": true}

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

			_, err = db.UpdateWithKeys(conn, &v, prefix, keys)

			if err != nil {
				return nil, err
			}

			*item = v
		}
	}

	//清除缓存
	{
		cache, err := app.GetCache("default")

		if err == nil {

			cache.Del(fmt.Sprintf("%d_%d", task.Tid, iid))

		}
	}

	data := AddData{Code: EncodeToString(iids), Items: items}

	app.SendMessage(task.GetName(), &data)

	return &data, nil
}
