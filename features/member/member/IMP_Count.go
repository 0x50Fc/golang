package member

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

	data := CountData{}

	maxSecond := time.Duration(dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"cache", "maxSecond"}), 60))

	var key string

	if task.Bid != nil {
		key = fmt.Sprintf("B_%d", dynamic.IntValue(task.Bid, 0))
	} else if task.Uid != nil {
		key = fmt.Sprintf("U_%d", dynamic.IntValue(task.Uid, 0))
	} else {
		key = "Q"
	}

	cacheKey := cache.SignKey("n", task.Bid, task.Uid, task.Q)

	//从缓存中读取
	{
		cache, err := app.GetCache("default")
		if err == nil {
			text, err := cache.GetItem(key, cacheKey, time.Second*maxSecond)
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

	sql := bytes.NewBuffer(nil)

	args := []interface{}{}

	sql.WriteString(" WHERE 1")

	if task.Bid != nil {
		sql.WriteString(" AND bid=?")
		args = append(args, task.Bid)
	}

	if task.Uid != nil {
		sql.WriteString(" AND uid=?")
		args = append(args, task.Uid)
	}

	if task.Q != nil {
		q := fmt.Sprintf("%%%s%%", dynamic.StringValue(task.Q, ""))
		sql.WriteString(" AND (title LIKE ? OR keyword LIKE ?)")
		args = append(args, q, q)
	}

	v := Member{}

	n, err := db.Count(conn, &v, prefix, sql.String(), args...)

	if err != nil {
		return nil, err
	}

	data.Total = int32(n)

	{
		cache, err := app.GetCache("default")

		if err == nil {

			b, _ := json.Marshal(&data)

			cache.SetItem(key, cacheKey, string(b), time.Second*maxSecond)

		}
	}

	return &data, nil
}
