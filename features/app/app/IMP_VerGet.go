package app

import (
	"fmt"
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) VerGet(app micro.IContext, task *VerGetTask) (*Ver, error) {

	v := Ver{}

	expires := time.Duration(dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"cache", "maxSecond"}), 1800)) * time.Second

	cache, _ := app.GetCache("default")

	if cache != nil {

		text, err := cache.GetItem(fmt.Sprintf("%d", task.Appid), fmt.Sprintf("%d", task.Ver), expires)

		if err == nil {
			err = json.Unmarshal([]byte(dynamic.StringValue(text, "")), &v)
			if err == nil {
				return &v, nil
			}
		}
	}

	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}

	prefix = Prefix(app, prefix, task.Appid)

	p, err := db.Get(conn, &v, prefix, "WHERE appid=? AND ver=?", task.Appid, task.Ver)

	if p == nil {
		return nil, micro.NewError(ERROR_NOT_FOUND, "未找到App版本")
	}

	if cache != nil {
		b, _ := json.Marshal(&v)
		cache.SetItem(fmt.Sprintf("%d", task.Appid), fmt.Sprintf("%d", task.Ver), string(b), expires)
	}

	return &v, nil
}
