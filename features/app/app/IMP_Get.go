package app

import (
	"bytes"
	"fmt"
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Get(app micro.IContext, task *GetTask) (*App, error) {

	v := App{}

	expires := time.Duration(dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"cache", "maxSecond"}), 1800)) * time.Second

	cache, _ := app.GetCache("default")

	if cache != nil {

		text, err := cache.Get(fmt.Sprintf("%d", task.Id), expires)

		if err == nil {
			err = json.Unmarshal([]byte(dynamic.StringValue(text, "")), &v)
			if err == nil {
				if task.Uid != nil {
					if dynamic.IntValue(task.Uid, 0) != v.Uid {
						return nil, micro.NewError(ERROR_NOT_FOUND, "未找到App")
					}
				}
				return &v, nil
			}
		}
	}

	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}

	prefix = Prefix(app, prefix, task.Id)

	sql := bytes.NewBuffer(nil)
	args := []interface{}{}

	sql.WriteString(" WHERE id=?")
	args = append(args, task.Id)

	if task.Uid != nil {
		sql.WriteString(" AND uid=?")
		args = append(args, task.Uid)
	}

	p, err := db.Get(conn, &v, prefix, sql.String(), args...)

	if p == nil {
		return nil, micro.NewError(ERROR_NOT_FOUND, "未找到App")
	}

	if cache != nil {
		b, _ := json.Marshal(&v)
		cache.Set(fmt.Sprintf("%d", task.Id), string(b), expires)
	}

	return &v, nil
}
