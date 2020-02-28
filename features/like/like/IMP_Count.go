package like

import (
	"bytes"
	"fmt"
	"time"

	"github.com/hailongz/golang/cache"
	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Count(app micro.IContext, task *CountTask) (*CountData, error) {

	iid := dynamic.IntValue(task.Iid, 0)

	data := CountData{}

	maxSecond := time.Duration(dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"cache", "maxSecond"}), 60))

	cacheKey := cache.SignKey("n", iid, task.Uid)

	//从缓存中读取
	{
		cache, err := app.GetCache("default")
		if err == nil {
			text, err := cache.GetItem(fmt.Sprintf("%d", task.Tid), cacheKey, time.Second*maxSecond)
			if err == nil && text != "" {
				err = json.Unmarshal([]byte(text), &data)
				if err == nil {
					return &data, nil
				}
			}
		}
	}

	//缓存没有查询数据库
	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}

	prefix = Prefix(app, prefix, task.Tid)

	sql := bytes.NewBuffer(nil)

	args := []interface{}{}

	sql.WriteString(" WHERE tid=? AND iid=?")

	args = append(args, task.Tid, iid)

	if task.Uid != nil {
		sql.WriteString(" AND uid=?")
		args = append(args, task.Uid)
	}

	v := Like{}

	count, err := db.Count(conn, &v, prefix, sql.String(), args...)

	if err != nil {
		return nil, err
	}

	data.Total = int32(count)

	{
		cache, err := app.GetCache("default")

		if err == nil {

			b, _ := json.Marshal(&data)

			cache.SetItem(fmt.Sprintf("%d", task.Tid), cacheKey, string(b), time.Second*maxSecond)

		}
	}

	return &data, nil
}
